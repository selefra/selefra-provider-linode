package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeVolumeGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeVolumeGenerator{}

func (x *TableLinodeVolumeGenerator) GetTableName() string {
	return "linode_volume"
}

func (x *TableLinodeVolumeGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeVolumeGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeVolumeGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeVolumeGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := linodego.ListOptions{}

			items, err := conn.ListVolumes(ctx, &opts)
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

func (x *TableLinodeVolumeGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeVolumeGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("The Volume’s size, in GiB.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("When this Volume was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filesystem_path").ColumnType(schema.ColumnTypeString).Description("The full filesystem path for the Volume based on the Volume’s label.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("Region where the volume resides.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated").ColumnType(schema.ColumnTypeTimestamp).Description("When this Volume was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("The unique ID of this Volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("The Volume’s label is for display purposes only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("linode_id").ColumnType(schema.ColumnTypeInt).Description("If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The current status of the volume: creating, active, resizing, contact_support.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Tags applied to this volume as a map.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("List of Tags applied to this volume.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableLinodeVolumeGenerator) GetSubTables() []*schema.Table {
	return nil
}
