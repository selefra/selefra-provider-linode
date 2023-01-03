package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeInstanceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeInstanceGenerator{}

func (x *TableLinodeInstanceGenerator) GetTableName() string {
	return "linode_instance"
}

func (x *TableLinodeInstanceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeInstanceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeInstanceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeInstanceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := linodego.ListOptions{}

			items, err := conn.ListInstances(ctx, &opts)
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

func (x *TableLinodeInstanceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeInstanceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("backups").ColumnType(schema.ColumnTypeJSON).Description("Information about this Linode’s backups status.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image").ColumnType(schema.ColumnTypeString).Description("An Image ID to deploy the Disk from.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6").ColumnType(schema.ColumnTypeCIDR).Description("This Linode’s IPv6 SLAAC address.").
			Extractor(column_value_extractor.StructSelector("IPv6")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The current status of the instance: creating, active, resizing, contact_support.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("The Instance’s label is for display purposes only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("List of Tags applied to this instance.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("watchdog_enabled").ColumnType(schema.ColumnTypeBool).Description("The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv4").ColumnType(schema.ColumnTypeJSON).Description("Array of this Linode’s IPv4 Addresses.").
			Extractor(column_value_extractor.StructSelector("IPv4")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("Region where the instance resides.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Tags applied to this instance as a map.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("specs").ColumnType(schema.ColumnTypeJSON).Description("Information about the resources available to this Linode, e.g. disk space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated").ColumnType(schema.ColumnTypeTimestamp).Description("When this Instance was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("The unique ID of this Instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alerts").ColumnType(schema.ColumnTypeJSON).Description("Alerts are triggered if CPU, IO, etc exceed these limits.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("When this Instance was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hypervisor").ColumnType(schema.ColumnTypeString).Description("The virtualization software powering this Linode, e.g. kvm.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).Description("This is the Linode Type that this Linode was deployed with.").
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
	}
}

func (x *TableLinodeInstanceGenerator) GetSubTables() []*schema.Table {
	return nil
}
