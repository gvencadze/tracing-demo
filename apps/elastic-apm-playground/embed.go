package apm_playground

import "embed"

//go:embed migrations/*.sql
var EmbedMigrations embed.FS
