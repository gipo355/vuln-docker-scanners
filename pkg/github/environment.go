//nolint:var-naming // we want to keep the same name as the GitHub Actions environment variables
package github

import (
	"fmt"
	"os"
)

// Variable	Description
//
// GitHub removes special characters, and uses the name __run when the current step runs a script without an id. If you use the same script or action more than once in the same job, the name will include a suffix that consists of the sequence number preceded by an underscore. For example, the first script you run will have the name __run, and the second script will be named __run_2. Similarly, the second invocation of actions/checkout will be actionscheckout2.

type Environment struct {
	// HOME The path to the home directory of the runner. For example, /home/runner.
	HOME string
	// GITHUB_WORKSPACE	The default working directory on the runner for steps, and the default location of your repository when using the checkout action. For example, /home/runner/work/my-repo-name/my-repo-name.
	// This is mounted as a volume on GitHub-hosted runners, and is the CWD of the runner.
	// If using actions/checkout, the working directory is automatically set to $GITHUB_WORKSPACE.
	// Writing to this folder allows persisted data to be passed between steps in a workflow run and allows accessing repo files.
	GITHUB_WORKSPACE string
	// CI	Always set to true.
	CI string
	// GITHUB_OUTPUT	The path on the runner to the file that sets the current step's outputs from workflow commands. The path to this file is unique to the current step and changes for each step in a job. For example, /home/runner/work/_temp/_runner_file_commands/set_output_a50ef383-b063-46d9-9157-57953fc9f3f0. For more information, see "Workflow commands for GitHub Actions."
	GITHUB_OUTPUT string
	// GITHUB_ENV	The path on the runner to the file that sets variables from workflow commands. The path to this file is unique to the current step and changes for each step in a job. For example, /home/runner/work/_temp/_runner_file_commands/set_env_87406d6e-4979-4d42-98e1-3dab1f48b13a. For more information, see "Workflow commands for GitHub Actions."
	GITHUB_ENV string
	// GITHUB_ACTION	The name of the action currently running, or the id of a step. For example, for an action, __repo-owner_name-of-action-repo.
	GITHUB_ACTION string
	// GITHUB_ACTION_PATH	The path where an action is located. This property is only supported in composite actions. You can use this path to change directories to where the action is located and access other files in that same repository. For example, /home/runner/work/_actions/repo-owner/name-of-action-repo/v1.
	GITHUB_ACTION_PATH string
	// GITHUB_ACTION_REPOSITORY	For a step executing an action, this is the owner and repository name of the action. For example, actions/checkout.
	GITHUB_ACTION_REPOSITORY string
	// GITHUB_ACTOR	The name of the person or app that initiated the workflow. For example, octocat.
	GITHUB_ACTOR string
	// GITHUB_ACTOR_ID	The account ID of the person or app that triggered the initial workflow run. For example, 1234567. Note that this is different from the actor username.
	GITHUB_ACTOR_ID string
	// GITHUB_ACTIONS	Always set to true when GitHub Actions is running the workflow. You can use this variable to differentiate when tests are being run locally or by GitHub Actions.
	GITHUB_ACTIONS string
	// GITHUB_API_URL	Returns the API URL. For example: https://api.github.com.
	GITHUB_API_URL string
	// GITHUB_BASE_REF	The name of the base ref or target branch of the pull request in a workflow run. This is only set when the event that triggers a workflow run is either pull_request or pull_request_target. For example, main.
	GITHUB_BASE_REF string
	// GITHUB_EVENT_NAME	The name of the event that triggered the workflow. For example, workflow_dispatch.
	GITHUB_EVENT_NAME string
	// GITHUB_EVENT_PATH	The path to the file on the runner that contains the full event webhook payload. For example, /github/workflow/event.json.
	GITHUB_EVENT_PATH string
	// GITHUB_GRAPHQL_URL	Returns the GraphQL API URL. For example: https://api.github.com/graphql.
	GITHUB_GRAPHQL_URL string
	// GITHUB_HEAD_REF	The head ref or source branch of the pull request in a workflow run. This property is only set when the event that triggers a workflow run is either pull_request or pull_request_target. For example, feature-branch-1.
	GITHUB_HEAD_REF string
	// GITHUB_JOB	The job_id of the current job. For example, greeting_job.
	GITHUB_JOB string
	// GITHUB_PATH	The path on the runner to the file that sets system PATH variables from workflow commands. The path to this file is unique to the current step and changes for each step in a job. For example, /home/runner/work/_temp/_runner_file_commands/add_path_899b9445-ad4a-400c-aa89-249f18632cf5. For more information, see "Workflow commands for GitHub Actions."
	GITHUB_PATH string
	// GITHUB_REF	The fully-formed ref of the branch or tag that triggered the workflow run. For workflows triggered by push, this is the branch or tag ref that was pushed. For workflows triggered by pull_request, this is the pull request merge branch. For workflows triggered by release, this is the release tag created. For other triggers, this is the branch or tag ref that triggered the workflow run. This is only set if a branch or tag is available for the event type. The ref given is fully-formed, meaning that for branches the format is refs/heads/<branch_name>, for pull requests it is refs/pull/<pr_number>/merge, and for tags it is refs/tags/<tag_name>. For example, refs/heads/feature-branch-1.
	GITHUB_REF string
	// GITHUB_REF_NAME	The short ref name of the branch or tag that triggered the workflow run. This value matches the branch or tag name shown on GitHub. For example, feature-branch-1.
	// For pull requests, the format is <pr_number>/merge.
	GITHUB_REF_NAME string
	// GITHUB_REF_PROTECTED	true if branch protections or rulesets are configured for the ref that triggered the workflow run.
	GITHUB_REF_PROTECTED string
	// GITHUB_REF_TYPE	The type of ref that triggered the workflow run. Valid values are branch or tag.
	GITHUB_REF_TYPE string
	// GITHUB_REPOSITORY	The owner and repository name. For example, octocat/Hello-World.
	GITHUB_REPOSITORY string
	// GITHUB_REPOSITORY_ID	The ID of the repository. For example, 123456789. Note that this is different from the repository name.
	GITHUB_REPOSITORY_ID string
	// GITHUB_REPOSITORY_OWNER	The repository owner's name. For example, octocat.
	GITHUB_REPOSITORY_OWNER string
	// GITHUB_REPOSITORY_OWNER_ID	The repository owner's account ID. For example, 1234567. Note that this is different from the owner's name.
	GITHUB_REPOSITORY_OWNER_ID string
	// GITHUB_RETENTION_DAYS	The number of days that workflow run logs and artifacts are kept. For example, 90.
	GITHUB_RETENTION_DAYS string
	// GITHUB_RUN_ATTEMPT	A unique number for each attempt of a particular workflow run in a repository. This number begins at 1 for the workflow run's first attempt, and increments with each re-run. For example, 3.
	GITHUB_RUN_ATTEMPT string
	// GITHUB_SERVER_URL	The URL of the GitHub server. For example: https://github.com.
	GITHUB_SERVER_URL string
	// GITHUB_SHA	The commit SHA that triggered the workflow. The value of this commit SHA depends on the event that triggered the workflow. For more information, see "Events that trigger workflows." For example, ffac537e6cbbf934b08745a378932722df287a53.
	GITHUB_SHA string
	// GITHUB_RUN_ID	A unique number for each workflow run within a repository. This number does not change if you re-run the workflow run. For example, 1658821493.
	GITHUB_RUN_ID string
	// GITHUB_RUN_NUMBER	A unique number for each run of a particular workflow in a repository. This number begins at 1 for the workflow's first run, and increments with each new run. This number does not change if you re-run the workflow run. For example, 3.
	GITHUB_RUN_NUMBER string
	// GITHUB_STEP_SUMMARY	The path on the runner to the file that contains job summaries from workflow commands. The path to this file is unique to the current step and changes for each step in a job. For example, /home/runner/_layout/_work/_temp/_runner_file_commands/step_summary_1cb22d7f-5663-41a8-9ffc-13472605c76c. For more information, see "Workflow commands for GitHub Actions."
	GITHUB_STEP_SUMMARY string
	// GITHUB_TRIGGERING_ACTOR	The username of the user that initiated the workflow run. If the workflow run is a re-run, this value may differ from github.actor. Any workflow re-runs will use the privileges of github.actor, even if the actor initiating the re-run (github.triggering_actor) has different privileges.
	GITHUB_TRIGGERING_ACTOR string
	// GITHUB_WORKFLOW	The name of the workflow. For example, My test workflow. If the workflow file doesn't specify a name, the value of this variable is the full path of the workflow file in the repository.
	GITHUB_WORKFLOW string
	// GITHUB_WORKFLOW_REF	The ref path to the workflow. For example, octocat/hello-world/.github/workflows/my-workflow.yml@refs/heads/my_branch.
	GITHUB_WORKFLOW_REF string
	// GITHUB_WORKFLOW_SHA	The commit SHA for the workflow file.
	GITHUB_WORKFLOW_SHA string
	// RUNNER_ARCH	The architecture of the runner executing the job. Possible values are X86, X64, ARM, or ARM64.
	RUNNER_ARCH string
	// RUNNER_DEBUG	This is set only if debug logging is enabled, and always has the value of 1. It can be useful as an indicator to enable additional debugging or verbose logging in your own job steps.
	GITHUB_RUNNER_DEBUG string
	// RUNNER_NAME	The name of the runner executing the job. This name may not be unique in a workflow run as runners at the repository and organization levels could use the same name. For example, Hosted Agent
	GITHUB_RUNNER_NAME string
	// RUNNER_OS	The operating system of the runner executing the job. Possible values are Linux, Windows, or macOS. For example, Windows
	RUNNER_OS string

	// RUNNER_TEMP	The path to a temporary directory on the runner. This directory is emptied at the beginning and end of each job. Note that files will not be removed if the runner's user account does not have permission to delete them. For example, D:\a\_temp
	RUNNER_TEMP string
	// RUNNER_TOOL_CACHE	The path to the directory containing preinstalled tools for GitHub-hosted runners. For more information, see "Using GitHub-hosted runners". For example, C:\hostedtoolcache\windows
	RUNNER_TOOL_CACHE string

	// GITHUB_STATE	The path on the runner to the file that sets the current step's state from workflow commands. The path to this file is unique to the current step and changes for each step in a job. For example, /home/runner/work/_temp/_runner_file_commands/set_state_1cb22d7f-5663-41a8-9ffc-13472605c76c. For more information, see "Workflow commands for GitHub Actions."
	GITHUB_STATE string
}

func NewGitHubEnvironment() (*Environment, error) {
	if os.Getenv("CI") != "true" {
		return nil, fmt.Errorf("not running in CI")
	}

	return &Environment{
		HOME:                       os.Getenv("HOME"),
		GITHUB_WORKSPACE:           os.Getenv("GITHUB_WORKSPACE"),
		CI:                         os.Getenv("CI"),
		GITHUB_OUTPUT:              os.Getenv("GITHUB_OUTPUT"),
		GITHUB_ENV:                 os.Getenv("GITHUB_ENV"),
		GITHUB_ACTION:              os.Getenv("GITHUB_ACTION"),
		GITHUB_ACTION_PATH:         os.Getenv("GITHUB_ACTION_PATH"),
		GITHUB_ACTION_REPOSITORY:   os.Getenv("GITHUB_ACTION_REPOSITORY"),
		GITHUB_ACTOR:               os.Getenv("GITHUB_ACTOR"),
		GITHUB_ACTOR_ID:            os.Getenv("GITHUB_ACTOR_ID"),
		GITHUB_ACTIONS:             os.Getenv("GITHUB_ACTIONS"),
		GITHUB_API_URL:             os.Getenv("GITHUB_API_URL"),
		GITHUB_BASE_REF:            os.Getenv("GITHUB_BASE_REF"),
		GITHUB_EVENT_NAME:          os.Getenv("GITHUB_EVENT_NAME"),
		GITHUB_EVENT_PATH:          os.Getenv("GITHUB_EVENT_PATH"),
		GITHUB_GRAPHQL_URL:         os.Getenv("GITHUB_GRAPHQL_URL"),
		GITHUB_HEAD_REF:            os.Getenv("GITHUB_HEAD_REF"),
		GITHUB_JOB:                 os.Getenv("GITHUB_JOB"),
		GITHUB_PATH:                os.Getenv("GITHUB_PATH"),
		GITHUB_REF:                 os.Getenv("GITHUB_REF"),
		GITHUB_REF_NAME:            os.Getenv("GITHUB_REF_NAME"),
		GITHUB_REF_PROTECTED:       os.Getenv("GITHUB_REF_PROTECTED"),
		GITHUB_REF_TYPE:            os.Getenv("GITHUB_REF_TYPE"),
		GITHUB_REPOSITORY:          os.Getenv("GITHUB_REPOSITORY"),
		GITHUB_REPOSITORY_ID:       os.Getenv("GITHUB_REPOSITORY_ID"),
		GITHUB_REPOSITORY_OWNER:    os.Getenv("GITHUB_REPOSITORY_OWNER"),
		GITHUB_REPOSITORY_OWNER_ID: os.Getenv("GITHUB_REPOSITORY_OWNER_ID"),
		GITHUB_RETENTION_DAYS:      os.Getenv("GITHUB_RETENTION_DAYS"),
		GITHUB_RUN_ATTEMPT:         os.Getenv("GITHUB_RUN_ATTEMPT"),
		GITHUB_SERVER_URL:          os.Getenv("GITHUB_SERVER_URL"),
		GITHUB_SHA:                 os.Getenv("GITHUB_SHA"),
		GITHUB_RUN_ID:              os.Getenv("GITHUB_RUN_ID"),
		GITHUB_RUN_NUMBER:          os.Getenv("GITHUB_RUN_NUMBER"),
		GITHUB_STEP_SUMMARY:        os.Getenv("GITHUB_STEP_SUMMARY"),
		GITHUB_TRIGGERING_ACTOR:    os.Getenv("GITHUB_TRIGGERING_ACTOR"),
		GITHUB_WORKFLOW:            os.Getenv("GITHUB_WORKFLOW"),
		GITHUB_WORKFLOW_REF:        os.Getenv("GITHUB_WORKFLOW_REF"),
		GITHUB_WORKFLOW_SHA:        os.Getenv("GITHUB_WORKFLOW_SHA"),
		RUNNER_ARCH:                os.Getenv("RUNNER_ARCH"),
		GITHUB_RUNNER_DEBUG:        os.Getenv("GITHUB_RUNNER_DEBUG"),
		GITHUB_RUNNER_NAME:         os.Getenv("GITHUB_RUNNER_NAME"),
		RUNNER_OS:                  os.Getenv("RUNNER_OS"),
		RUNNER_TEMP:                os.Getenv("RUNNER_TEMP"),
		RUNNER_TOOL_CACHE:          os.Getenv("RUNNER_TOOL_CACHE"),
		GITHUB_STATE:               os.Getenv("GITHUB_STATE"),
	}, nil
}
