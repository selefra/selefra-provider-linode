package tables

import (
	"context"
	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeDomainRecordGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeDomainRecordGenerator{}

func (x *TableLinodeDomainRecordGenerator) GetTableName() string {
	return "linode_domain_record"
}

func (x *TableLinodeDomainRecordGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeDomainRecordGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeDomainRecordGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeDomainRecordGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := linodego.ListOptions{}

			domainData := task.ParentRawResult.(*linodego.Domain)
			domainID := domainData.ID

			var items []linodego.DomainRecord
			if opts.Filter == "" {

				items, err = conn.ListDomainRecords(ctx, domainID, nil)
			} else {
				items, err = conn.ListDomainRecords(ctx, domainID, &opts)
			}
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

func (x *TableLinodeDomainRecordGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeDomainRecordGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("This Record’s unique ID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeInt).Description("The priority of the target host for this Record. Lower values are preferred. Only valid for MX and SRV record requests. Required for SRV record requests.").
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protocol").ColumnType(schema.ColumnTypeString).Description("The protocol this Record’s service communicates with. An underscore (_) is prepended automatically to the submitted value for this property. Only valid for SRV record requests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service").ColumnType(schema.ColumnTypeString).Description("The name of the service. Only valid and required for SRV record requests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag").ColumnType(schema.ColumnTypeString).Description("The tag portion of a CAA record. Only valid and required for CAA record requests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("weight").ColumnType(schema.ColumnTypeInt).Description("The relative weight of this Record used in the case of identical priority. Higher values are preferred. Only valid and required for SRV record requests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("record_type").ColumnType(schema.ColumnTypeString).Description("The type of Record this is in the DNS system: A, AAAA, NS, MX, CNAME, TXT, SRV, PTR, CAA.").
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of this Record. This property’s actual usage and whether it is required depends on the type of record it represents. For example, for CNAME, it is the hostname.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeInt).Description("The port this Record points to. Only valid and required for SRV record requests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeString).Description("The target for this Record. For requests, this property’s actual usage and whether it is required depends on the type of record this represents. For example, for CNAME it is the domain target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl_sec").ColumnType(schema.ColumnTypeInt).Description("Time to Live - the amount of time in seconds that the domain record may be cached by resolvers or other domain servers.").
			Extractor(column_value_extractor.StructSelector("TTLSec")).Build(),
	}
}

func (x *TableLinodeDomainRecordGenerator) GetSubTables() []*schema.Table {
	return nil
}
