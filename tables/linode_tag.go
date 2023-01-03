package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableLinodeTagGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeTagGenerator{}

func (x *TableLinodeTagGenerator) GetTableName() string {
	return "linode_tag"
}

func (x *TableLinodeTagGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeTagGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeTagGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeTagGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			items, err := conn.ListTags(ctx, nil)
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

func (x *TableLinodeTagGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeTagGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("A Label used for organization of objects on your Account.").Build(),
	}
}

func (x *TableLinodeTagGenerator) GetSubTables() []*schema.Table {
	return nil
}
