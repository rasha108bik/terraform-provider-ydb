package index

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tbl "github.com/ydb-platform/terraform-provider-ydb/internal/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func (h *handler) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	indexResource, err := indexResourceSchemaToIndexResource(d)
	if err != nil {
		return diag.FromErr(err)
	}

	db, err := tbl.CreateDBConnection(ctx, tbl.ClientParams{
		DatabaseEndpoint: indexResource.ConnectionString,
		Token:            h.token,
	})
	if err != nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "failed to initialize table client",
				Detail:   err.Error(),
			},
		}
	}
	defer func() {
		_ = db.Close(ctx)
	}()

	q := prepareCreateIndexRequest(indexResource)
	err = db.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		return s.ExecuteSchemeQuery(ctx, q)
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(indexResource.ConnectionString + "/" + indexResource.TablePath + "/" + indexResource.Name)

	return h.Read(ctx, d, meta)
}
