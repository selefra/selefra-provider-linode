package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableLinodeRegionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeRegionGenerator{}

func (x *TableLinodeRegionGenerator) GetTableName() string {
	return "linode_region"
}

func (x *TableLinodeRegionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeRegionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeRegionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeRegionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			items, err := conn.ListRegions(ctx, nil)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range items {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableLinodeRegionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeRegionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country").ColumnType(schema.ColumnTypeString).Description("Country for the region.").Build(),
	}
}

func (x *TableLinodeRegionGenerator) GetSubTables() []*schema.Table {
	return nil
}
