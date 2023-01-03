package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableLinodeTokenGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeTokenGenerator{}

func (x *TableLinodeTokenGenerator) GetTableName() string {
	return "linode_token"
}

func (x *TableLinodeTokenGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeTokenGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeTokenGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeTokenGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			items, err := conn.ListTokens(ctx, nil)
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

func (x *TableLinodeTokenGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeTokenGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("This token's label. This is for display purposes only, but can be used to more easily track what you're using each token for.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("token").ColumnType(schema.ColumnTypeString).Description("First 16 characters of the token.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scopes").ColumnType(schema.ColumnTypeJSON).Description("Array of scopes for the token, e.g. *, account:read_write, domains:read_only.").
			Extractor(column_value_extractor.StructSelector("Scopes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time this token was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiry").ColumnType(schema.ColumnTypeTimestamp).Description("When this token will expire.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("This token's unique ID, which can be used to revoke it.").Build(),
	}
}

func (x *TableLinodeTokenGenerator) GetSubTables() []*schema.Table {
	return nil
}
