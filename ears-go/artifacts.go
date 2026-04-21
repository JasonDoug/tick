package main

type Feature struct {
	ID          string `json:"id" description:"Short ID (e.g., F-001)"`
	Name        string `json:"name" description:"Short colloquial name"`
	Description string `json:"description" description:"1-2 sentence user-friendly description"`
}

type Scope struct {
	InScope    []string `json:"in_scope" description:"List of in-scope items"`
	OutOfScope []string `json:"out_of_scope" description:"List of out-of-scope items"`
}

type OnePager struct {
	DocumentType       string    `json:"document_type" description:"Must be 'Project_Charter_OnePager'"`
	NorthStar          string    `json:"northstar" description:"One concise sentence unifying the vision"`
	Features           []Feature `json:"features" description:"3-8 MVP features"`
	SuccessMetrics     []string  `json:"success_metrics" description:"Measurable success criteria"`
	Scope              Scope     `json:"scope" description:"In-scope and out-of-scope boundaries"`
	NonFunctionalNotes string    `json:"non_functional_notes" description:"High-level constraints"`
	Status             string    `json:"status" description:"'draft' or 'approved'"`
	Version            string    `json:"version" description:"Document version"`
}

type Requirement struct {
	ID            string `json:"id" description:"Requirement ID (e.g., R-F001-001)"`
	FeatureID     string `json:"feature_id" description:"ID of the feature it relates to"`
	EARSStatement string `json:"ears_statement" description:"Requirement in EARS format (e.g., 'When..., the system shall...')"`
	Type          string `json:"type" description:"'functional' or 'non-functional'"`
	Priority      string `json:"priority" description:"'must_have', 'should_have', or 'could_have'"`
}

type UseCase struct {
	ID                   string   `json:"id" description:"Use case ID (e.g., UC-001)"`
	Name                 string   `json:"name" description:"Name of the use case"`
	RealizesRequirements []string `json:"realizes_requirements" description:"IDs of requirements this use case implements"`
	Preconditions        string   `json:"preconditions" description:"Required state before the use case"`
	MainFlow             []string `json:"main_flow" description:"Step-by-step main successful path"`
	AlternateFlows       []string `json:"alternate_flows" description:"Alternative paths"`
	Postconditions       string   `json:"postconditions" description:"State after successful completion"`
	Exceptions           []string `json:"exceptions" description:"Failure conditions"`
}

type Architecture struct {
	TechStack  []string `json:"tech_stack" description:"List of technologies used"`
	Modules    []string `json:"modules" description:"Major system components"`
	DataModels []string `json:"data_models" description:"Key data structures"`
}

type ADD struct {
	DocumentType string        `json:"document_type" description:"Must be 'Architectural_Design_Document'"`
	Version      string        `json:"version" description:"Document version"`
	Requirements []Requirement `json:"requirements" description:"List of EARS requirements"`
	UseCases     []UseCase     `json:"use_cases" description:"List of use cases"`
	Architecture Architecture  `json:"architecture" description:"System architecture definition"`
	Status       string        `json:"status" description:"'draft' or 'approved'"`
}

type SubTask struct {
	ID               string `json:"id" description:"Sub-task ID (e.g., ST-001-1)"`
	Prompt           string `json:"prompt" description:"Focused prompt for the AI coding tool"`
	DefinitionOfDone string `json:"definition_of_done" description:"Criteria for completion"`
}

type Task struct {
	ID       string    `json:"id" description:"Task ID (e.g., T-001)"`
	Name     string    `json:"name" description:"Short name of the task"`
	SubTasks []SubTask `json:"sub_tasks" description:"Atomic coding prompts"`
}

type Sprint struct {
	ID    string `json:"id" description:"Sprint ID (e.g., S-001)"`
	Name  string `json:"name" description:"Name of the sprint"`
	Tasks []Task `json:"tasks" description:"List of tasks in the sprint"`
}

type ProjectPlan struct {
	DocumentType string   `json:"document_type" description:"Must be 'Project_Plan'"`
	Version      string   `json:"version" description:"Document version"`
	Sprints      []Sprint `json:"sprints" description:"List of logical work groupings"`
}
