// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/graphql/generated"
	"github.com/facebookincubator/symphony/graph/graphql/resolver"
	"github.com/facebookincubator/symphony/graph/viewer"
	"github.com/facebookincubator/symphony/pkg/log"
	"github.com/facebookincubator/symphony/pkg/mysql"

	"go.uber.org/zap"
)

type cliFlags struct {
	Dsn    string `env:"MYSQL_DSN" long:"dsn" description:"data source name"`
	Tenant string `long:"tenant" required:"true" description:"target specific tenant"`
	User   string `long:"user" required:"true" description:"target specific user"`
}

func main() {
	logger, _ := log.Config{Format: "console"}.Build()
	ctx := context.Background()

	var cf cliFlags
	if _, err := flags.Parse(&cf); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	logger.For(ctx).Info("params", zap.String("dsn", cf.Dsn), zap.String("tenant", cf.Tenant), zap.String("user", cf.User))

	tenancy, err := viewer.NewMySQLTenancy(cf.Dsn)
	if err != nil {
		logger.For(ctx).Fatal("cannot connect to graph database", zap.String("dsn", cf.Dsn), zap.Error(err))
		return
	}

	mysql.SetLogger(logger)

	v := &viewer.Viewer{Tenant: cf.Tenant, User: cf.User}

	ctx = log.NewFieldsContext(ctx, zap.Object("viewer", v))
	ctx = viewer.NewContext(ctx, v)

	client, err := tenancy.ClientFor(ctx, cf.Tenant)
	if err != nil {
		logger.For(ctx).Fatal("cannot get ent client for tenant", zap.String("tenant", cf.Tenant), zap.Error(err))
		return
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		logger.For(ctx).Error("cannot begin transaction", zap.Error(err))
		return
	}

	defer func() {
		if r := recover(); r != nil {
			if err := tx.Rollback(); err != nil {
				logger.For(ctx).Error("cannot rollback transaction", zap.Error(err))
			}
			panic(r)
		}
	}()

	ctx = ent.NewContext(ctx, tx.Client())

	// Since the client is already uses transaction we can't have transactions on graphql also
	r, err := resolver.New(logger, resolver.WithTransaction(false))
	if err != nil {
		logger.For(ctx).Error("cannot initialize graphql resolver", zap.Error(err))
		return
	}

	if err := utilityFunc(ctx, r, logger); err != nil {
		logger.For(ctx).Error("failed to run function", zap.Error(err))
		if err := tx.Rollback(); err != nil {
			logger.For(ctx).Error("cannot rollback transaction", zap.Error(err))
		}
		return
	}

	if err := tx.Commit(); err != nil {
		logger.For(ctx).Error("cannot commit transaction", zap.Error(err))
	}
}

func utilityFunc(ctx context.Context, r generated.ResolverRoot, logger log.Logger) error {
	/**
	Add your Go code in this function
	You need to run this code from the same version production is at to avoid schema mismatches
	DO NOT LAND THE CODE AFTER THIS COMMENT
	*/
	/*
		Example code:
		client := ent.FromContext(ctx)
		eqt, err := r.Mutation().AddEquipmentType(ctx, models.AddEquipmentTypeInput{Name: "My new type"})
		if err != nil {
			return errors.Wrap(err, "failed to create equipment type")
		}
		logger.For(ctx).Info("equipment created", zap.String("ID", eqt.ID))
		client.EquipmentType.UpdateOneID(eqt.ID).SetName("My new type 2").ExecX(ctx)
		if err != nil {
			return errors.Wrapf(err, "failed to update equipment type id=%q", eqt.ID)
		}
	*/
	return nil
}
