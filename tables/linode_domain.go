package tables

import (
	"context"
	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/linode_client"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeDomainGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeDomainGenerator{}

func (x *TableLinodeDomainGenerator) GetTableName() string {
	return "linode_domain"
}

func (x *TableLinodeDomainGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeDomainGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeDomainGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeDomainGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := linodego.ListOptions{}

			items, err := conn.ListDomains(ctx, &opts)
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

func (x *TableLinodeDomainGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeDomainGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("soa_email").ColumnType(schema.ColumnTypeString).Description("Start of Authority email address. This is required for master Domains.").
			Extractor(column_value_extractor.StructSelector("SOAEmail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Used to control whether this Domain is currently being rendered: disabled, active, edit_mode, has_errors.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain").ColumnType(schema.ColumnTypeString).Description("The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_ips").ColumnType(schema.ColumnTypeJSON).Description("The IP addresses representing the master DNS for this Domain.").
			Extractor(column_value_extractor.StructSelector("MasterIPs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A description for this Domain. This is for display purposes only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl_sec").ColumnType(schema.ColumnTypeInt).Description("Time to Live - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers.").
			Extractor(column_value_extractor.StructSelector("TTLSec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_type").ColumnType(schema.ColumnTypeString).Description("If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).").
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retry_sec").ColumnType(schema.ColumnTypeInt).Description("The interval, in seconds, at which a failed refresh should be retried.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Description("List of Tags applied to this domain.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("The unique ID of this Domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("axfr_ips").ColumnType(schema.ColumnTypeJSON).Description("The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.").
			Extractor(column_value_extractor.StructSelector("AXfrIPs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Tags applied to this domain as a map.").
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_sec").ColumnType(schema.ColumnTypeInt).Description("The amount of time in seconds that may pass before this Domain is no longer authoritative.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("refresh_sec").ColumnType(schema.ColumnTypeInt).Description("The amount of time in seconds before this Domain should be refreshed.").Build(),
	}
}

func (x *TableLinodeDomainGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableLinodeDomainRecordGenerator{}),
	}
}
