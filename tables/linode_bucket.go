package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableLinodeBucketGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeBucketGenerator{}

func (x *TableLinodeBucketGenerator) GetTableName() string {
	return "linode_bucket"
}

func (x *TableLinodeBucketGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeBucketGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeBucketGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeBucketGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			items, err := conn.ListObjectStorageBuckets(ctx, nil)
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

func (x *TableLinodeBucketGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeBucketGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("The name of this bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster").ColumnType(schema.ColumnTypeString).Description("The ID of the Object Storage Cluster this bucket is in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("When this bucket was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hostname").ColumnType(schema.ColumnTypeString).Description("The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public.").Build(),
	}
}

func (x *TableLinodeBucketGenerator) GetSubTables() []*schema.Table {
	return nil
}
