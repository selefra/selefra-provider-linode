package tables

import (
	"context"

	"github.com/selefra/selefra-provider-linode/linode_client"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableLinodeAccountGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeAccountGenerator{}

func (x *TableLinodeAccountGenerator) GetTableName() string {
	return "linode_account"
}

func (x *TableLinodeAccountGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeAccountGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeAccountGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeAccountGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			item, err := conn.GetAccount(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- *item
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableLinodeAccountGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeAccountGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("address_2").ColumnType(schema.ColumnTypeString).Description("Second line of this Account’s billing address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("city").ColumnType(schema.ColumnTypeString).Description("The city for this Account’s billing address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("company").ColumnType(schema.ColumnTypeString).Description("The company name associated with this Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_name").ColumnType(schema.ColumnTypeString).Description("The last name of the person associated with this Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone").ColumnType(schema.ColumnTypeString).Description("The phone number associated with this Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The email address of the person associated with this Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address_1").ColumnType(schema.ColumnTypeString).Description("First line of this Account’s billing address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("balance_uninvoiced").ColumnType(schema.ColumnTypeString).Description("This Account’s current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country").ColumnType(schema.ColumnTypeString).Description("The two-letter country code of this Account’s billing address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("The state for this Account’s billing address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tax_id").ColumnType(schema.ColumnTypeString).Description("The tax identification number associated with this Account, for tax calculations in some countries. If you do not live in a country that collects tax, this should be null.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("balance").ColumnType(schema.ColumnTypeString).Description("This Account’s balance, in US dollars.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credit_card").ColumnType(schema.ColumnTypeJSON).Description("Credit Card information associated with this Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_name").ColumnType(schema.ColumnTypeString).Description("The first name of the person associated with this Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zip").ColumnType(schema.ColumnTypeString).Description("The zip code of this Account’s billing address.").Build(),
	}
}

func (x *TableLinodeAccountGenerator) GetSubTables() []*schema.Table {
	return nil
}
