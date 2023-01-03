package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableLinodeUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeUserGenerator{}

func (x *TableLinodeUserGenerator) GetTableName() string {
	return "linode_user"
}

func (x *TableLinodeUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := linodego.ListOptions{}

			items, err := conn.ListUsers(ctx, &opts)
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

func (x *TableLinodeUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).Description("This User’s username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The email address for this User, for account management communications, and may be used for other communications as configured.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restricted").ColumnType(schema.ColumnTypeBool).Description("If true, this User must be granted access to perform actions or access entities on this Account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssh_keys").ColumnType(schema.ColumnTypeJSON).Description("A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request.").Build(),
	}
}

func (x *TableLinodeUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
