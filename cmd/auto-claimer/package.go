package main

import (
	"github.com/SlashNephy/auto-claimer/config"
	"github.com/SlashNephy/auto-claimer/database"
	"github.com/SlashNephy/auto-claimer/pipeline"
	"github.com/SlashNephy/auto-claimer/query"
	"github.com/SlashNephy/auto-claimer/repository"
	"github.com/SlashNephy/auto-claimer/workflow"
	"github.com/samber/do/v2"
)

var Package = do.Package(
	config.Package,
	database.Package,
	repository.Package,
	query.Package,
	workflow.Package,
	pipeline.Package,
)
