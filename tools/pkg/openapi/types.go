// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package openapi provides types and utilities for parsing OpenAPI specifications
// for the F5XC Terraform provider code generation tools.
package openapi

// Spec represents an OpenAPI 3.x specification.
type Spec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       map[string]interface{} `json:"info"`
	Paths      map[string]interface{} `json:"paths"`
	Components Components             `json:"components"`
}

// Components contains the reusable components of an OpenAPI spec.
type Components struct {
	Schemas map[string]Schema `json:"schemas"`
}

// RequiredFor describes which operations a field is required for.
type RequiredFor struct {
	MinimumConfig bool `json:"minimum_config"`
	Create        bool `json:"create"`
	Update        bool `json:"update"`
	Read          bool `json:"read"`
}

// ConcurrencyToken describes a server-assigned optimistic-concurrency field.
// The field is carried by the API client and Terraform private state, never by
// the practitioner-facing resource schema.
type ConcurrencyToken struct {
	ServerAssigned   bool     `json:"server_assigned"`
	EchoOnOperations []string `json:"echo_on_operations"`
}

// Schema represents a schema definition in an OpenAPI spec.
type Schema struct {
	Type                 string            `json:"type"`
	Description          string            `json:"description"`
	Title                string            `json:"title"`
	Format               string            `json:"format"`
	Enum                 []interface{}     `json:"enum"`
	Default              interface{}       `json:"default"`
	ReadOnly             bool              `json:"readOnly"`
	Properties           map[string]Schema `json:"properties"`
	Items                *Schema           `json:"items"`
	Ref                  string            `json:"$ref"`
	Required             []string          `json:"required"`
	AdditionalProperties interface{}       `json:"additionalProperties"`
	// AllOf wraps $ref for OAS3 compliance (x-ves-* sibling preservation pattern)
	AllOf []Schema `json:"allOf"`

	// Original F5 vendor extensions (x-ves-*) - technical metadata from upstream
	XDisplayName        string            `json:"x-displayname"`
	XVesExample         string            `json:"x-ves-example"`
	XVesValidationRules map[string]string `json:"x-ves-validation-rules"`
	XVesProtoMessage    string            `json:"x-ves-proto-message"`

	// Enrichment extensions (x-f5xc-*) - added by api-specs-enriched repository
	XF5XCCategory         string                `json:"x-f5xc-category"`
	XF5XCRequiresTier     string                `json:"x-f5xc-requires-tier"`
	XF5XCRequires         []RequiresEntry       `json:"x-f5xc-requires"`
	XF5XCComplexity       string                `json:"x-f5xc-complexity"`
	XF5XCExample          string                `json:"x-f5xc-example"`
	XF5XCDescriptionShort string                `json:"x-f5xc-description-short"`
	XF5XCDescriptionMed   string                `json:"x-f5xc-description-medium"`
	XF5XCUseCases         []string              `json:"x-f5xc-use-cases"`
	XF5XCRelatedDomains   []string              `json:"x-f5xc-related-domains"`
	XF5XCIsPreview        bool                  `json:"x-f5xc-is-preview"`
	XF5XCNamespaceProfile *NamespaceProfileSpec `json:"x-f5xc-namespace-profile"`
	XF5XCIcon             string                `json:"x-f5xc-icon"`
	XF5XCCLIDomain        string                `json:"x-f5xc-cli-domain"`

	// Additional upstream extensions
	XVesDeprecated   string            `json:"x-ves-deprecated"`
	XVesDisplayOrder string            `json:"x-ves-displayorder"`
	XVesProtoEnum    string            `json:"x-ves-proto-enum"`
	XVesRequired     string            `json:"x-ves-required"`
	XRequired        bool              `json:"x-required"`
	XValidationRules map[string]string `json:"x-validation-rules"`

	// Enrichment — actionable in generation
	XF5XCConflictsWith           []string               `json:"x-f5xc-conflicts-with"`
	XF5XCConstraints             map[string]interface{} `json:"x-f5xc-constraints"`
	XF5XCRecommendedOneofVariant interface{}            `json:"x-f5xc-recommended-oneof-variant"`
	XFieldMutability             string                 `json:"x-field-mutability"`
	XOriginalMaxLength           int                    `json:"x-original-maxLength"`

	// Provenance — parsed, not used in generation
	XReconciledAt            string `json:"x-reconciled-at"`
	XReconciledFromDiscovery bool   `json:"x-reconciled-from-discovery"`

	// ---- SP-1 additions: field-level extensions ----
	XF5XCServerDefault        bool              `json:"x-f5xc-server-default"`
	XF5XCRequiredFor          RequiredFor       `json:"x-f5xc-required-for"`
	XF5XCRecommendedValue     interface{}       `json:"x-f5xc-recommended-value"`
	XF5XCMinimumConfiguration interface{}       `json:"x-f5xc-minimum-configuration"`
	XF5XCConcurrencyToken     *ConcurrencyToken `json:"x-f5xc-concurrency-token,omitempty"`

	// Deserialized from enriched specs for lossless round-tripping.
	// No generation code consumes these yet.
	XF5XCValidation        map[string]interface{} `json:"x-f5xc-validation"`
	XF5XCDefaults          map[string]interface{} `json:"x-f5xc-defaults"`
	XF5XCConditions        map[string]interface{} `json:"x-f5xc-conditions"`
	XF5XCDeprecated        string                 `json:"x-f5xc-deprecated"`
	XF5XCCompletion        map[string]interface{} `json:"x-f5xc-completion"`
	XF5XCDisplayName       string                 `json:"x-f5xc-display-name"`
	XF5XCDescription       string                 `json:"x-f5xc-description"`
	XF5XCExamples          []interface{}          `json:"x-f5xc-examples"`
	XF5XCRequiredForOps    map[string]interface{} `json:"x-f5xc-required-for-operations"`
	XF5XCDisplayOrder      int                    `json:"x-f5xc-displayorder"`
	XF5XCUniqueness        string                 `json:"x-f5xc-uniqueness"`
	XF5XCTerraformResource string                 `json:"x-f5xc-terraform-resource"`

	// ---- SP-1 additions: operation-level extensions ----
	XF5XCDangerLevel          string                 `json:"x-f5xc-danger-level"`
	XF5XCConfirmationRequired bool                   `json:"x-f5xc-confirmation-required"`
	XF5XCSideEffects          []string               `json:"x-f5xc-side-effects"`
	XF5XCOperationMetadata    map[string]interface{} `json:"x-f5xc-operation-metadata"`
	XF5XCRequiredFields       []string               `json:"x-f5xc-required-fields"`

	// XF5xcAction marks a request-body schema whose operation is an action-style
	// resource (e.g. "approve") so codegen emits an action resource — Create=action
	// POST, Read=sibling object Get, Delete=no-op — instead of CRUD. Injected by
	// api-specs-enriched onto the request schema (schema-level, not operation-level).
	XF5xcAction string `json:"x-f5xc-action"`

	// XF5xcWireName records the property's ON-THE-WIRE JSON key when it differs
	// from the property name. F5 ships several misspelled property names
	// (blocked_sevice, public_advertisment, disable_lb_source_ip_persistance) and
	// the wire key must stay misspelled — verified live: a PUT with
	// `blocked_sevice` returns HTTP 200 and round-trips, while `blocked_service`
	// is silently ignored by the server (#1257). The Terraform attribute name is
	// ours to choose, so api-specs-enriched presents the CORRECTED spelling as the
	// property name and records the original key here. Codegen derives the
	// Terraform attribute, model field and docs from the property name, and uses
	// this wire name for every marshal AND unmarshal key. Absent, the property
	// name is the wire key, so the properties without it are unchanged. See #1323.
	XF5xcWireName string `json:"x-f5xc-wire-name"`
}

// Int64RangeSpan is one inclusive interval in a possibly discontinuous integer
// constraint. Multiple spans are ordered and non-overlapping in generator IR.
type Int64RangeSpan struct {
	Minimum int64
	Maximum int64
}

// TerraformAttribute represents an attribute in a Terraform resource schema.
type TerraformAttribute struct {
	Name               string
	GoName             string
	TfsdkTag           string
	Type               string
	ElementType        string
	Description        string
	Required           bool
	CreateRequired     bool // Required when the containing optional block is configured.
	Optional           bool
	Computed           bool
	Sensitive          bool
	NestedAttributes   []TerraformAttribute
	NestedBlockType    string
	IsBlock            bool
	ConversionError    string // Tracks controlled generator errors during conversion
	OneOfGroup         string
	PlanModifier       string
	MaxDepth           int    // Track recursion depth to prevent infinite loops
	IsSpecField        bool   // True if this is a spec field (not metadata)
	JsonName           string // On-the-wire JSON key for API marshaling AND unmarshaling: Schema.WireName (x-f5xc-wire-name, else the property name). Never the Terraform name — see TfsdkTag.
	GoType             string // Go type for client struct generation
	UseDomainValidator bool   // True if name field should use DomainValidator (for DNS resources)

	// ---- SP-1 additions: enrichment-driven attribute metadata ----
	ServerDefault         bool              // x-f5xc-server-default
	Default               interface{}       // Resolved default value
	MinimumConfigRequired bool              // Derived from x-f5xc-required-for.minimum_config
	RecommendedValue      interface{}       // x-f5xc-recommended-value
	ValidationRules       map[string]string // Merged x-ves-validation-rules + x-validation-rules
	Complexity            string            // x-f5xc-complexity
	UseCases              []string          // x-f5xc-use-cases
	DeprecationMessage    string            // x-f5xc-deprecated or x-ves-deprecated
	ConflictsWith         []string          // x-f5xc-conflicts-with
	MaxLength             int               // From x-original-maxLength or validation rules
	Immutable             bool              // x-field-mutability == "immutable"
	EnumValues            []string          // Resolved from enum + x-ves-proto-enum
	StringDefault         string            // Spec-driven static string default (e.g. namespace fixed to "system"); empty = none
	MinLength             int               // From validation rules
	Pattern               string            // From validation rules
	Format                string            // From x-f5xc-constraints.format — drives string format validators (ipv4/ipv6/ip/cidr/mac-address)
	ETLDPlusOne           bool              // From ves.io.schema.rules.string.etld_plus_one — value must be an eTLD+1 domain
	MinItems              int               // From x-f5xc-constraints.min_items
	MaxItems              int               // From x-f5xc-constraints.max_items
	Minimum               int
	Maximum               int
	// HasMinimum/HasMaximum record presence of the numeric bound (independent of
	// value), so an int64 field with minimum:0 still emits a range validator.
	HasMinimum bool
	HasMaximum bool
	// Int64RangeSpans preserves range-set rules such as "0,512-16384" without
	// widening the gap into a single minimum/maximum interval.
	Int64RangeSpans []Int64RangeSpan
}

// ResourceTemplate contains data for generating a Terraform resource.
type ResourceTemplate struct {
	Name                   string
	TitleCase              string
	APIPath                string
	APIPathPlural          string
	APIPathItem            string // Path for single item operations (get/update/delete)
	HasNamespaceInPath     bool   // Whether API path contains namespace segment
	Description            string
	Attributes             []TerraformAttribute
	OneOfGroups            map[string][]string
	HasComplexSpec         bool
	RequiredAttributes     []string
	OptionalAttributes     []string
	ComputedAttributes     []string
	ExampleUsage           string // HCL example for documentation
	APIDocsURL             string // Link to F5 XC API documentation
	UsesBoolPlanModifier   bool   // True if any bool attribute uses a plan modifier
	UsesInt64PlanModifier  bool   // True if any int64 attribute uses a plan modifier
	UsesStringPlanModifier bool   // True if any string attribute uses a plan modifier
	UsesListPlanModifier   bool   // True if any list attribute uses a plan modifier
	UsesMapPlanModifier    bool   // True if any map attribute uses a plan modifier

	// ---- SP-1 additions: generation control flags ----
	HasBlocks               bool // True if any attribute is a block
	HasMaxLengthValidators  bool // True if any attribute has MaxLength > 0
	HasEnumValidators       bool // True if any attribute has EnumValues
	HasPatternValidators    bool // True if any attribute has Pattern
	HasListSizeValidators   bool // True if any attribute has MinItems or MaxItems
	HasInt64RangeValidators bool
	HasStringDefaults       bool   // True if any attribute has a StringDefault (needs stringdefault import)
	HasConflicts            bool   // True if any attribute has ConflictsWith
	ConflictCheckCode       string // Generated Go code for conflict checks
	IsReadOnly              bool   // True if resource has GetSpecType only (data source, no resource)

	// Preflights are apply-time requirement checks compiled into Create/Update.
	// Source of truth: x-f5xc-requires (api-specs-enriched). Loaded from
	// tools/preflight-requirements.json keyed by TitleCase; see preflight.go.
	Preflights []RequirementPreflight

	// ImportIDExtraFields are create-only spec fields the API cannot return on read,
	// so their values ride in the import ID as trailing segments after namespace/name.
	// Loaded from tools/import-id-fields.json keyed by TitleCase; see import_id_fields.go.
	ImportIDExtraFields []string

	// ExposeUID makes the generator surface the object's server-generated
	// system_metadata.uid as a Computed, read-only `uid` attribute, carry it on the
	// client struct (SystemMetadata), and populate it from the Create/Get response.
	// It is opt-in per resource (tools/expose-uid.json, keyed by TitleCase) AND
	// gated on the response schema actually carrying system_metadata.uid, so it
	// stays surgically scoped and never emits a uid a resource cannot return.
	// See tools/pkg/openapi/expose_uid.go and schema.ResponseHasSystemMetadataUID.
	ExposeUID bool

	// FiltersDiscoveredSiteLabels makes the generated Read drop the six hardware/OS
	// discovery labels F5 XC writes onto a node-backed site (domain, host-os-version,
	// hw-model, hw-serial-number, hw-vendor, hw-version) instead of letting them into
	// state, where an empty `labels` block proposed deleting them on every plan
	// (#1391). Opt-in per resource (tools/discovered-site-labels.json, keyed by
	// TitleCase) because those names are ordinary enough that a user could own one on
	// an unrelated object. The two platform prefixes are filtered regardless.
	// See tools/pkg/openapi/discovered_site_labels.go.
	FiltersDiscoveredSiteLabels bool

	// Concurrency-token fields are detected from matching GET response and replace
	// request envelopes. They intentionally do not appear in Attributes.
	HasConcurrencyToken      bool
	ConcurrencyTokenJSONName string
	ConcurrencyTokenGoName   string
	// UpdateDisabled marks either an API with no Replace operation or an
	// evidence-backed non-config-object Replace exclusion. Every practitioner-
	// settable field then requires replacement, and Update fails closed without
	// sending a PUT if the framework invokes it unexpectedly.
	UpdateDisabled bool

	// ---- Action-resource fields (x-f5xc-action) ----
	// IsAction marks this template as an action-style resource: Create issues the
	// action POST (ActionPath), Read does a lenient GET on the sibling object
	// (ReadObjectPath) with 404 -> remove-from-state, Delete is a no-op, and there
	// is no Update and no data-source companion. ActionState is the constant state
	// value the action drives the object to (e.g. "APPROVED").
	IsAction       bool
	ActionPath     string // %s-substituted singular action POST path
	ActionState    string // constant state the action applies (e.g. "APPROVED")
	ReadObjectPath string // %s-substituted sibling object GET path (pluralized)

	// ActionDerivedFields are request-body fields the action REQUIRES but that
	// are facts about the object being acted on rather than user input. The
	// generated Create reads the sibling object (ReadObjectPath, leniently) and
	// echoes each value back verbatim. Loaded from tools/action-derived-fields.json.
	ActionDerivedFields []ActionDerivedField
}

// ActionDerivedField declares one server-derived field of an action request
// body: a value the API demands but will only accept as the exact object it
// already holds, so it can never be a Terraform attribute.
//
// The F5 XC registration approve is the motivating case (#1355): the POST is
// rejected with HTTP 500 "Validation approval: Passport is required" unless it
// carries the registration's own spec.gc_spec.passport, and the API accepts only
// that object echoed back. Exposing it as an attribute would let a practitioner
// supply a passport the server refuses; omitting it made the resource incapable
// of ever creating. Reading it off the object closes both.
//
// Sources are dotted lookup paths into the leniently parsed sibling read,
// evaluated in order with the first hit winning: the read wraps the same value
// at more than one depth (object.spec.gc_spec.passport is the canonical object
// shape, spec.passport the flattened projection), and which wrappers appear
// varies with the endpoint.
type ActionDerivedField struct {
	Field    string   `json:"field"`   // request-body property name, e.g. "passport"
	Sources  []string `json:"sources"` // dotted paths into the sibling read; first hit wins
	GoName   string   `json:"-"`       // request-struct field name, e.g. "Passport" (derived)
	JSONName string   `json:"-"`       // wire key for the request body (derived)
}

// RequirementPreflight is one apply-time prerequisite the provider verifies
// before writing a resource. When the model field WhenGoField is set (non-nil),
// the generated Create/Update lists ListPath in the resource's namespace and, if
// the collection is empty, fails fast with ErrorTitle/ErrorDetail — turning an
// opaque server error into an actionable remediation. It encodes, in the shipped
// binary, the dependency declared by x-f5xc-requires (e.g. client_side_defense
// requires a same-namespace protected_domain), so every remote workstation
// enforces it identically without relying on out-of-band knowledge.
type RequirementPreflight struct {
	WhenField   string `json:"when_field"`   // JSON/tfsdk name of the triggering field, e.g. "client_side_defense"
	WhenGoField string `json:"-"`            // Go model field to nil-check, e.g. "ClientSideDefense" (resolved from attributes)
	ListPath    string `json:"list_path"`    // LIST path with a single %s for the namespace
	Requires    string `json:"requires"`     // Human-readable requirement (rendered as a code comment)
	ErrorTitle  string `json:"error_title"`  // Diagnostic summary
	ErrorDetail string `json:"error_detail"` // Diagnostic detail; must contain exactly one %s for the namespace
}

// GenerationResult tracks the result of generating a resource.
type GenerationResult struct {
	ResourceName string
	Success      bool
	Error        string
	FilePath     string
	// Attributes carries the generated Terraform schema IR for validations that
	// must run after extraction (for example SMSv2 legacy parity).
	Attributes []TerraformAttribute

	// ---- SP-1 additions: generation metrics ----
	AttrCount  int  // Number of attributes generated
	BlockCount int  // Number of nested blocks generated
	IsReadOnly bool // True if only GetSpecType found (data source only, no resource)
	IsAction   bool // True if this is an action-style resource (x-f5xc-action)
}

// IsRef returns true if the schema is a reference to another schema.
func (s *Schema) IsRef() bool {
	return s.Ref != ""
}

// IsArray returns true if the schema type is array.
func (s *Schema) IsArray() bool {
	return s.Type == "array"
}

// IsObject returns true if the schema type is object.
func (s *Schema) IsObject() bool {
	return s.Type == "object"
}

// IsRequired checks if a property name is in the required list.
func (s *Schema) IsRequired(propertyName string) bool {
	for _, r := range s.Required {
		if r == propertyName {
			return true
		}
	}
	return false
}

// HasProperties returns true if the schema has properties defined.
func (s *Schema) HasProperties() bool {
	return len(s.Properties) > 0
}

// WireName returns the JSON key this property must use ON THE WIRE: the
// x-f5xc-wire-name override when present, otherwise propName itself.
//
// It exists so codegen can split the two names that used to be one. The
// Terraform attribute, model field and docs come from propName (which the buffer
// zone spells correctly); every marshal and unmarshal key comes from this, so a
// property F5 misspells keeps round-tripping. Both sides must use it: if only the
// request used the wire key, the read-back would look up a key the API never
// returns and the field would silently drift. See XF5xcWireName and #1323.
//
// The annotation is a fact about the PROPERTY, not about a component it $refs, so
// callers must read it from the property schema before resolving any $ref — a
// shared component must not rename every property that references it.
func (s *Schema) WireName(propName string) string {
	if s.XF5xcWireName != "" {
		return s.XF5xcWireName
	}
	return propName
}

// =============================================================================
// V2 Spec Types - For parsing enriched API specifications from api-specs-enriched
// =============================================================================

// Index represents the index.json manifest file in v2 spec structure.
// This file provides metadata about all domain specifications.
type Index struct {
	Version           string                 `json:"version"`
	Timestamp         string                 `json:"timestamp"`
	Specifications    []DomainMetadata       `json:"specifications"`
	CriticalResources []string               `json:"x-f5xc-critical-resources"`
	ErrorResolution   map[string]interface{} `json:"x-f5xc-error-resolution"`
	GuidedWorkflows   map[string]interface{} `json:"x-f5xc-guided-workflows"`
	Acronyms          map[string]interface{} `json:"x-f5xc-acronyms"`
}

// DomainMetadata represents metadata about a domain specification file.
// Field names map to the x-f5xc-* extensions in index.json.
type DomainMetadata struct {
	Name              string                    `json:"domain"` // Domain name from "domain" field
	File              string                    `json:"file"`
	Category          string                    `json:"x-f5xc-category"`
	Description       string                    `json:"description"`
	DescriptionShort  string                    `json:"x-f5xc-description-short"`
	DescriptionMedium string                    `json:"x-f5xc-description-medium"`
	Icon              string                    `json:"x-f5xc-icon"`
	RequiresTier      string                    `json:"x-f5xc-requires-tier"`
	Complexity        string                    `json:"x-f5xc-complexity"`
	IsPreview         bool                      `json:"x-f5xc-is-preview"`
	CLIDomain         string                    `json:"x-f5xc-cli-domain"`
	Title             string                    `json:"title"`
	PathCount         int                       `json:"path_count"`
	SchemaCount       int                       `json:"schema_count"`
	RelatedDomains    []string                  `json:"x-f5xc-related-domains"`
	UseCases          []string                  `json:"x-f5xc-use-cases"`
	PrimaryResources  []PrimaryResourceMetadata `json:"x-f5xc-primary-resources"`
}

// PrimaryResourceMetadata represents resource-level metadata from x-f5xc-primary-resources in index.json.
// This is extracted from index.json and provides per-resource tier and dependency info.
type PrimaryResourceMetadata struct {
	Name              string               `json:"name"`
	Description       string               `json:"description"`
	DescriptionShort  string               `json:"description_short"`
	Tier              string               `json:"tier"`
	Icon              string               `json:"icon"`
	Category          string               `json:"category"`
	SupportsLogs      bool                 `json:"supports_logs"`
	SupportsMetrics   bool                 `json:"supports_metrics"`
	Dependencies      ResourceDependencies `json:"dependencies"`
	RelationshipHints []string             `json:"relationship_hints"`

	// ---- SP-1 additions: schema and API path references ----
	SchemaComponents []string `json:"schema_components"`
	APIPaths         []string `json:"api_paths"`
}

// ResourceDependencies represents the dependencies of a resource.
type ResourceDependencies struct {
	Required []string `json:"required"`
	Optional []string `json:"optional"`
}

// ResourceMetadata represents metadata about a resource within a domain.
// Used by ExtractResourcesFromDomain for processing.
type ResourceMetadata struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	APIPath              string   `json:"api_path"`
	RequiresTier         string   `json:"requires_tier"`
	Complexity           string   `json:"complexity"`
	Dependencies         []string `json:"dependencies"`
	MinimumConfiguration string   `json:"minimum_configuration"`
}

// DomainSpec represents a parsed domain specification file (v2 format).
// Each domain file contains multiple related resources.
type DomainSpec struct {
	OpenAPI    string                 `json:"openapi"`
	Info       DomainInfo             `json:"info"`
	Paths      map[string]interface{} `json:"paths"`
	Components Components             `json:"components"`

	// Domain-level enrichment metadata
	XF5XCCategory         string                `json:"x-f5xc-category"`
	XF5XCRequiresTier     string                `json:"x-f5xc-requires-tier"`
	XF5XCComplexity       string                `json:"x-f5xc-complexity"`
	XF5XCIsPreview        bool                  `json:"x-f5xc-is-preview"`
	XF5XCRelatedDomains   []string              `json:"x-f5xc-related-domains"`
	XF5XCUseCases         []string              `json:"x-f5xc-use-cases"`
	XF5XCNamespaceProfile *NamespaceProfileSpec `json:"x-f5xc-namespace-profile"`

	// ---- SP-1 additions: domain-level spec metadata ----
	XF5XCDocSection        string                    `json:"x-f5xc-doc-section"`
	XF5XCPrimaryResources  []PrimaryResourceMetadata `json:"x-f5xc-primary-resources"`
	XF5XCCriticalResources []string                  `json:"x-f5xc-critical-resources"`
	XF5XCLogoSVG           string                    `json:"x-f5xc-logo-svg"`
	XF5XCIcon              string                    `json:"x-f5xc-icon"`
}

// DomainInfo represents the info section of a domain spec.
type DomainInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`

	// Enrichment extensions at info level
	XF5XCDescriptionShort  string                `json:"x-f5xc-description-short"`
	XF5XCDescriptionMedium string                `json:"x-f5xc-description-medium"`
	XF5XCIcon              string                `json:"x-f5xc-icon"`
	XF5XCLogoSVG           string                `json:"x-f5xc-logo-svg"`
	XF5XCDescriptionLong   string                `json:"x-f5xc-description-long"`
	XF5XCSummary           string                `json:"x-f5xc-summary"`
	XF5XCBestPractices     *BestPractices        `json:"x-f5xc-best-practices"`
	XF5XCNamespaceProfile  *NamespaceProfileSpec `json:"x-f5xc-namespace-profile"`

	// ---- SP-1 additions: spec-level domain metadata ----
	XF5XCCLIMetadata     map[string]interface{} `json:"x-f5xc-cli-metadata"`
	XF5XCGlossary        map[string]interface{} `json:"x-f5xc-glossary"`
	XF5XCGuidedWorkflows []interface{}          `json:"x-f5xc-guided-workflows"`
	XF5XCAcronyms        map[string]interface{} `json:"x-f5xc-acronyms"`
}

// NamespaceProfileSpec represents the x-f5xc-namespace-profile extension object
// in an enriched OpenAPI spec. It captures namespace constraints, recommendations,
// and classification metadata for a resource.
type NamespaceProfileSpec struct {
	Constraint     *NamespaceConstraint     `json:"constraint,omitempty"`
	Recommendation *NamespaceRecommendation `json:"recommendation,omitempty"`
	Classification *NamespaceClassification `json:"classification,omitempty"`
}

// NamespaceConstraint defines which namespaces a resource is allowed in.
type NamespaceConstraint struct {
	Allowed  []string `json:"allowed"`
	Enforced bool     `json:"enforced"`
}

// NamespaceRecommendation defines the recommended namespace for a resource.
type NamespaceRecommendation struct {
	Primary string `json:"primary"`
}

// NamespaceClassification provides categorization metadata for a resource.
type NamespaceClassification struct {
	Category           string `json:"category"`
	MultiTenantPattern string `json:"multi_tenant_pattern"`
}

// BestPractices contains operational guidance from the enriched spec.
type BestPractices struct {
	CommonErrors []ErrorPattern `json:"common_errors"`
}

// ErrorPattern describes a common API error and how to resolve it.
type ErrorPattern struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Resolution string `json:"resolution"`
	Prevention string `json:"prevention"`
}

// SpecVersion represents the detected specification version.
type SpecVersion string

const (
	// SpecVersionV2 represents the v2 spec format (38 domain-organized files).
	SpecVersionV2 SpecVersion = "v2"
	// SpecVersionUnknown represents an unrecognized spec format.
	SpecVersionUnknown SpecVersion = "unknown"
)

// V2Categories maps domain categories from v2 specs.
var V2Categories = map[string]string{
	"security":       "Security",
	"networking":     "Networking",
	"infrastructure": "Infrastructure",
	"platform":       "Platform",
	"operations":     "Operations",
	"ai":             "AI Services",
}
