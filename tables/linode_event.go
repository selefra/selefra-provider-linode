package tables

import (
	"context"
	"github.com/selefra/selefra-provider-linode/linode_client"

	"github.com/linode/linodego"
	"github.com/selefra/selefra-provider-linode/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableLinodeEventGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableLinodeEventGenerator{}

func (x *TableLinodeEventGenerator) GetTableName() string {
	return "linode_event"
}

func (x *TableLinodeEventGenerator) GetTableDescription() string {
	return ""
}

func (x *TableLinodeEventGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableLinodeEventGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableLinodeEventGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := linode_client.Connect(ctx, taskClient.(*linode_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := linodego.ListOptions{}

			items, err := conn.ListEvents(ctx, &opts)
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

func (x *TableLinodeEventGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableLinodeEventGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("action").ColumnType(schema.ColumnTypeString).Description("The action that caused this Event. New actions may be added in the future.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Current status of the Event, Enum: 'failed' 'finished' 'notification' 'scheduled' 'started'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entity").ColumnType(schema.ColumnTypeJSON).Description("Detailed information about the Event's entity, including ID, type, label, and URL used to access it.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read").ColumnType(schema.ColumnTypeBool).Description("If this Event has been read.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time this event was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).Description("The username of the User who caused the Event.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("percent_complete").ColumnType(schema.ColumnTypeInt).Description("A percentage estimating the amount of time remaining for an Event. Returns null for notification events.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rate").ColumnType(schema.ColumnTypeString).Description("The rate of completion of the Event. Only some Events will return rate; for example, migration and resize Events.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("seen").ColumnType(schema.ColumnTypeBool).Description("If this Event has been seen.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_entity").ColumnType(schema.ColumnTypeJSON).Description("Detailed information about the Event's secondary or related entity, including ID, type, label, and URL used to access it.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_remaining").ColumnType(schema.ColumnTypeInt).Description("The estimated time remaining until the completion of this Event. This value is only returned for in-progress events.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("The unique ID of this Event.").Build(),
	}
}

func (x *TableLinodeEventGenerator) GetSubTables() []*schema.Table {
	return nil
}
