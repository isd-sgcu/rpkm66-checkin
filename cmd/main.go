package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/isd-sgcu/rpkm66-checkin/cfgldr"
	"github.com/isd-sgcu/rpkm66-checkin/database"
	event_v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/event/v1"
	namespace_v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/namespace/v1"
	staff_v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/staff/v1"
	event_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/event"
	namespace_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/namespace"
	staff_repo "github.com/isd-sgcu/rpkm66-checkin/pkg/repository/staff"
	event_service "github.com/isd-sgcu/rpkm66-checkin/pkg/service/event"
	namespace_service "github.com/isd-sgcu/rpkm66-checkin/pkg/service/namespace"
	staff_service "github.com/isd-sgcu/rpkm66-checkin/pkg/service/staff"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-s

		log.Info().
			Str("service", "graceful shutdown").
			Msgf("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Error().
				Str("service", "graceful shutdown").
				Msgf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Error().
						Str("service", "graceful shutdown").
						Err(err).
						Msgf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("%v was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}

func main() {
	conf, err := cfgldr.LoadConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
			Msg("Failed to start service")
	}

	db, err := database.InitDatabase(&conf.Database)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
			Msg("Failed to start service")
	}

	// cacheDB, err := database.InitRedisConnect(&conf.Redis)
	// if err != nil {
	// 	log.Fatal().
	// 		Err(err).
	// 		Str("service", "backend").
	// 		Msg("Failed to start service")
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
			Msg("Failed to start service")
	}

	namespaceRepo := namespace_repo.NewRepository(db)
	namespaceService := namespace_service.NewService(namespaceRepo)

	eventRepo := event_repo.NewRepository(db)
	eventService := event_service.NewService(eventRepo)

	staffRepo := staff_repo.NewRepository(db)
	staffService := staff_service.NewService(staffRepo)

	grpcServer := grpc.NewServer()

	// cacheRepo := cache.NewRepository(cacheDB)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	namespace_v1.RegisterNamespaceServiceServer(grpcServer, namespaceService)
	event_v1.RegisterEventServiceServer(grpcServer, eventService)
	staff_v1.RegisterStaffServiceServer(grpcServer, staffService)

	reflection.Register(grpcServer)
	go func() {
		log.Info().
			Str("service", "backend").
			Msgf("rpkm66 backend starting at port %v", conf.App.Port)

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal().
				Err(err).
				Str("service", "backend").
				Msg("Failed to start service")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			sqlDb, err := db.DB()
			if err != nil {
				return err
			}
			return sqlDb.Close()
		},
		// "cache": func(ctx context.Context) error {
		// 	return cacheDB.Close()
		// },
		"server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})

	<-wait

	grpcServer.GracefulStop()
	log.Info().
		Str("service", "backend").
		Msg("Closing the listener")
	lis.Close()
	log.Info().
		Str("service", "backend").
		Msg("End of Program")
}
