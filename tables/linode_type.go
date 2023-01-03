package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeTypeGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeTypeGenerator{}

func (x *TableLinodeTypeGenerator) GetTableName() string {
	return "linode_type"
}

func (x *TableLinodeTypeGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeTypeGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeTypeGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeTypeGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			items, err := conn.ListTypes(ctx, nil)
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

func (x *TableLinodeTypeGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeTypeGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("price_hourly").ColumnType(schema.ColumnTypeInt).Description("Cost (in US dollars) per hour.").
			Extractor(column_value_extractor.StructSelector("Price.Hourly")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("The Linode Type’s label is for display purposes only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addons").ColumnType(schema.ColumnTypeJSON).Description("A list of optional add-on services for Linodes and their associated costs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory").ColumnType(schema.ColumnTypeInt).Description("Amount of RAM included in this Linode Type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transfer").ColumnType(schema.ColumnTypeInt).Description("The monthly outbound transfer amount, in MB.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vcpus").ColumnType(schema.ColumnTypeInt).Description("The number of VCPU cores this Linode Type offers.").
			Extractor(column_value_extractor.StructSelector("VCPUs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID representing the Linode Type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk").ColumnType(schema.ColumnTypeInt).Description("The Disk size, in MB, of the Linode Type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_out").ColumnType(schema.ColumnTypeInt).Description("The Mbits outbound bandwidth allocation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("class").ColumnType(schema.ColumnTypeString).Description("The class of the Linode Type: nanode, standard, dedicated, gpu, highmem.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("price_monthly").ColumnType(schema.ColumnTypeInt).Description("Cost (in US dollars) per month.").
			Extractor(column_value_extractor.StructSelector("Price.Monthly")).Build(),
	}
}

func (x *TableLinodeTypeGenerator) GetSubTables() []*schema.Table {
	return nil
}
