// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-lamp/command"
)

const lampVersion string = "2.2.2"

var commonFlags = []gcli.Flag{
	gcli.BoolFlag{
		Name:  "v",
		Usage: "Execute commands in verbose mode",
	},
	gcli.StringFlag{
		Name:  "apiKey",
		Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
	},
	gcli.StringFlag{
		Name:  "user",
		Usage: "Owner of the execution",
	},
	gcli.StringFlag{
		Name:  "config",
		Usage: "Configuration file path",
	},
}

func createAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "message",
			Usage: "Alert text limited to 130 characters",
		},
		gcli.StringFlag{
			Name:  "recipients",
			Usage: "The user names of individual users or names of groups",
		},
		gcli.StringFlag{
			Name:  "teams",
			Usage: "A comma seperated list of teams",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "A user defined identifier for the alert and there can be only one alert with open status with the same alias.",
		},
		gcli.StringFlag{
			Name:  "actions",
			Usage: "A comma separated list of actions that can be executed",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Field to specify source of alert. By default, it will be assigned to IP address of incoming request",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "A comma separated list of labels attached to the alert",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Alert text in long form. Unlike the message field, not limited to 130 characters",
		},
		gcli.StringFlag{
			Name:  "entity",
			Usage: "The entity the alert is related to",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringSliceFlag{
			Name:  "D",
			Usage: "Additional alert properties.\n\tSyntax: -D key=value",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createAlert",
		Flags:  flags,
		Usage:  "Creates an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.CreateAlertAction(c)
			return nil
		},
	}
	return cmd
}

func getAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
		gcli.BoolFlag{
			Name:  "pretty",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getAlert",
		Flags:  flags,
		Usage:  "Gets an alert content from OpsGenie",
		Action: func(c *gcli.Context) error {
			command.GetAlertAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "createdAfter",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are created after specified time",
		},
		gcli.StringFlag{
			Name:  "createdBefore",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are created before specified time",
		},
		gcli.StringFlag{
			Name:  "updatedAfter",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are updated after specified time",
		},
		gcli.StringFlag{
			Name:  "updatedBefore",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are updated before specified time",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 20. Max value for this parameter is 100",
		},
		gcli.StringFlag{
			Name:  "status",
			Usage: "Used to query alerts with specified status. May take one of open, acked, unacked, seen, notseen, closed",
		},
		gcli.StringFlag{
			Name:  "sortBy",
			Usage: "createdAt, updatedAt, default is createdAt",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "asc/desc, default: desc",
		},
		gcli.StringFlag{
			Name:  "teams",
			Usage: "A comma seperated list of teams",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "A comma separated list of labels attached to the alert",
		},
		gcli.StringFlag{
			Name:  "tagsOperator",
			Usage: "tags are combined with tagsOperator when filtered. Accepted values: and/or, default: and",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
		gcli.BoolFlag{
			Name:  "pretty",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlerts",
		Flags:  flags,
		Usage:  "Lists alerts contents from OpsGenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertsAction(c)
			return nil
		},
	}
	return cmd
}

func countAlertsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "createdAfter",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are created after specified time",
		},
		gcli.StringFlag{
			Name:  "createdBefore",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are created before specified time",
		},
		gcli.StringFlag{
			Name:  "updatedAfter",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are updated after specified time",
		},
		gcli.StringFlag{
			Name:  "updatedBefore",
			Usage: "Unix timestamp value which is converted to nano second. Request will return all alerts which are updated before specified time",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 20. Max value for this parameter is 100",
		},
		gcli.StringFlag{
			Name:  "status",
			Usage: "Used to query alerts with specified status. May take one of open, acked, unacked, seen, notseen, closed",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "A comma separated list of labels attached to the alert",
		},
		gcli.StringFlag{
			Name:  "tagsOperator",
			Usage: "tags are combined with tagsOperator when filtered. Accepted values: and/or, default: and",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "countAlerts",
		Flags:  flags,
		Usage:  "Counts alerts at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.CountAlertsAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertNotesCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 100.",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "asc/desc, default : desc",
		},
		gcli.StringFlag{
			Name:  "lastKey",
			Usage: "Key which will be used in pagination.",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
		gcli.BoolFlag{
			Name:  "pretty",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlertNotes",
		Flags:  flags,
		Usage:  "Lists alert notes from OpsGenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertNotesAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertLogsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 100.",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "asc/desc, default : desc",
		},
		gcli.StringFlag{
			Name:  "lastKey",
			Usage: "Key which will be used in pagination.",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
		gcli.BoolFlag{
			Name:  "pretty",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlertLogs",
		Flags:  flags,
		Usage:  "Lists alert logs from OpsGenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertLogsAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertRecipientsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
		gcli.BoolFlag{
			Name:  "pretty",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlertRecipients",
		Flags:  flags,
		Usage:  "Lists alert recipients from OpsGenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertRecipientsAction(c)
			return nil
		},
	}
	return cmd
}

func unAcknowledgeCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be unacknowledged. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be unacknowledged. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "unacknowledge",
		Flags:  flags,
		Usage:  "Unacknowledges an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.UnAcknowledgeAction(c)
			return nil
		},
	}
	return cmd

}

func snoozeCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be snoozed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be snoozed. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "endDate",
			Usage: "The date and time snooze will end",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
		gcli.StringFlag{
			Name:  "timezone",
			Usage: "Timezone of endDate parameter",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "snooze",
		Flags:  flags,
		Usage:  "Snoozes an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.SnoozeAction(c)
			return nil
		},
	}
	return cmd

}

func removeTagsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the tags will be removed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the tags will be removed. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "A comma separated list of labels attached to the alert.",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "removeTags",
		Flags:  flags,
		Usage:  "Removes tags from an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.RemoveTagsAction(c)
			return nil
		},	}
	return cmd
}

func addDetailsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new details will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the new details will be added. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
		gcli.StringSliceFlag{
			Name:  "D",
			Usage: "Additional alert properties.\n\tSyntax: -D key=value",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addDetails",
		Flags:  flags,
		Usage:  "Adds details to an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.AddDetailsAction(c)
			return nil
		},	}
	return cmd
}

func removeDetailsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the details will be removed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the details will be removed. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "keys",
			Usage: "Set of properties to be removed from alert details",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "removeDetails",
		Flags:  flags,
		Usage:  "Removes details from an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.RemoveDetailsAction(c)
			return nil
		},	}
	return cmd
}

func escalateToNextActionCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the next escalation will be processed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the next escalation will be processed. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "escalationId",
			Usage: "Id of the escalation that will be escalated to the next level. Either escalationName or escalationId must be provided.",
		},
		gcli.StringFlag{
			Name:  "escalationName",
			Usage: "Name of the escalation that will be escalated to the next level. Either escalationName or escalationId must be provided.",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note text",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "escalateToNext",
		Flags:  flags,
		Usage:  "Esclates to the next rule in the specified escalation at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.EscalateToNextAction(c)
			return nil
		},	}
	return cmd
}

func attachFileCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the file will be attached. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the file will be attached. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "attachment",
			Usage: "Absolute or relative path to the file",
		},
		gcli.StringFlag{
			Name:  "indexFile",
			Usage: "",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "attachFile",
		Flags:  flags,
		Usage:  "Attaches files to an alert",
		Action: func(c *gcli.Context) error {
			command.AttachFileAction(c)
			return nil
		},
	}
	return cmd
}

func acknowledgeCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be acknowledged. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be acknowledged. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "acknowledge",
		Flags:  flags,
		Usage:  "Acknowledges an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.AcknowledgeAction(c)
			return nil
		},
	}
	return cmd

}

func renotifyCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that recipient will be renotified for. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that recipient will be renotified for. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "recipients",
			Usage: "The user names of individual users or names of groups that will be renotified for alert",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "renotify",
		Flags:  flags,
		Usage:  "Renotifies recipients at OpsGenie.",
		Action: func(c *gcli.Context) error {
			command.RenotifyAction(c)
			return nil
		},	}
	return cmd
}

func takeOwnershipCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be owned. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be owned. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "takeOwnership",
		Flags:  flags,
		Usage:  "Takes the ownership of an alert at OpsGenie.",
		Action: func(c *gcli.Context) error {
			command.TakeOwnershipAction(c)
			return nil
		},	}
	return cmd
}

func assignOwnerCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be owned. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be owned. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "owner",
			Usage: "The users who will be the owner of the alert after the execution",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "assign",
		Flags:  flags,
		Usage:  "Assigns the ownership of an alert to the specified user.",
		Action: func(c *gcli.Context) error {
			command.AssignOwnerAction(c)
			return nil
		},	}
	return cmd
}

func addTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new team will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the new team will be added. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "team",
			Usage: "The team that will be added to the alert",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addTeam",
		Flags:  flags,
		Usage:  "Adds a new team to an alert.",
		Action: func(c *gcli.Context) error {
			command.AddTeamAction(c)
			return nil
		},	}
	return cmd
}

func addRecipientCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new recipient will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the new recipient will be added. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "recipient",
			Usage: "The recipient that will be added to the alert",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addRecipient",
		Flags:  flags,
		Usage:  "Adds a new recipient to an alert.",
		Action: func(c *gcli.Context) error {
			command.AddRecipientAction(c)
			return nil
		},	}
	return cmd
}

func addNoteCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be retrieved. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note text",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addNote",
		Flags:  flags,
		Usage:  "Adds a user comment for an alert.",
		Action: func(c *gcli.Context) error {
			command.AddNoteAction(c)
			return nil
		},	}
	return cmd
}

func addTagsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new tags will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the new tags will be added. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "A comma separated list of labels attached to the alert.",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addTags",
		Flags:  flags,
		Usage:  "Adds tags to an alert.",
		Action: func(c *gcli.Context) error {
			command.AddTagsAction(c)
			return nil
		},	}
	return cmd
}

func executeActionCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the action will be executed on. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the action will be executed on. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "action",
			Usage: "Action to execute",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note text",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "executeAction",
		Flags:  flags,
		Usage:  "Executes alert actions at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.ExecuteActionAction(c)
			return nil
		},	}
	return cmd
}

func closeAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId,id",
			Usage: "Id of the alert that will be closed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be closed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "notify",
			Usage: "Comma separated list of user and groups which will be notified. Also special values \"all\", \"recipients\" and \"owner\" is accepted",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note text",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "closeAlert",
		Flags:  flags,
		Usage:  "Closes an alert at OpsGenie",
		Action: func(c *gcli.Context) error {
			command.CloseAlertAction(c)
			return nil
		},	}
	return cmd
}

func deleteAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be deleted",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteAlert",
		Flags:  flags,
		Usage:  "Deletes an alert at OpsGenie.",
		Action: func(c *gcli.Context) error {
			command.DeleteAlertAction(c)
			return nil
		},	}
	return cmd
}

func heartbeatCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the heartbeat on OpsGenie",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "heartbeat",
		Flags:  flags,
		Usage:  "Sends heartbeat to OpsGenie",
		Action: func(c *gcli.Context) error {
			command.HeartbeatAction(c)
			return nil
		},	}
	return cmd
}

func enableCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Id of the integration/policy that will be enabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the integration/policy that will be enabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "integration or policy",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "enable",
		Flags:  flags,
		Usage:  "Enables OpsGenie Integration and Policy.",
		Action: func(c *gcli.Context) error {
			command.EnableAction(c)
			return nil
		},	}
	return cmd
}

func disableCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Id of the integration/policy that will be disabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the integration/policy that will be disabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "integration or policy",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "disable",
		Flags:  flags,
		Usage:  "Disables OpsGenie Integration and Policy.",
		Action: func(c *gcli.Context) error {
			command.DisableAction(c)
			return nil
		},	}
	return cmd
}
func customerLogListDownloadablesCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "after",
			Usage: "Log files after this date",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listLogs",
		Flags:  flags,
		Usage:  "List Downloadable Customer Logs.",
		Action: func(c *gcli.Context) error {
			command.ListDownloadableLogs(c)
			return nil
		},	}
	return cmd
}
func customerLogGetLinkCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "logFile",
			Usage: "Buraya neyin logFile'i onun yazilmasi lazim",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getlogLink",
		Flags:  flags,
		Usage:  "Get link of log file",
		Action: func(c *gcli.Context) error {
			command.GetLogLink(c)
			return nil
		},	}
	return cmd
}

func initCommands(app *gcli.App) {
	app.Commands = []gcli.Command{
		createAlertCommand(),
		getAlertCommand(),
		attachFileCommand(),
		acknowledgeCommand(),
		renotifyCommand(),
		takeOwnershipCommand(),
		assignOwnerCommand(),
		addTeamCommand(),
		addRecipientCommand(),
		addTagsCommand(),
		addNoteCommand(),
		executeActionCommand(),
		closeAlertCommand(),
		deleteAlertCommand(),
		heartbeatCommand(),
		enableCommand(),
		disableCommand(),
		listAlertsCommand(),
		countAlertsCommand(),
		listAlertNotesCommand(),
		listAlertLogsCommand(),
		listAlertRecipientsCommand(),
		unAcknowledgeCommand(),
		snoozeCommand(),
		removeTagsCommand(),
		addDetailsCommand(),
		removeDetailsCommand(),
		escalateToNextActionCommand(),
		customerLogGetLinkCommand(),
		customerLogListDownloadablesCommand(),
	}
}

func main() {
	app := gcli.NewApp()
	app.Name = "lamp"
	app.Version = lampVersion
	app.Usage = "Command line interface for OpsGenie"
	app.Author = "OpsGenie"
	app.Action = func(c *gcli.Context) {
		fmt.Printf("Run 'lamp help' for the options\n")
	}
	initCommands(app)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error occured while executing command: %s\n", err.Error())
	}

}
