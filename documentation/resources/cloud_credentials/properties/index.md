---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_cloud_credentials."
xcsh_docs: {"aliases": ["authentication", "cloud credentials", "credential setup", "credentials"], "body_bytes": 24835, "body_sha256": "sha256:4889f5386af57d358fa38cda5cf5b5218dfc783427100aaf882560754fd74233", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key", "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file", "xcsh-docs:resources:cloud_credentials:properties:timeouts"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:reference", "parent_id": "xcsh-docs:resources:cloud_credentials:fundamentals", "path": "documentation/resources/cloud_credentials/properties/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3201213022001320-3300101112101223-0013231303001331-1203003101323223-3301132033321232-3133332231103012-2231201322220022-1313022312220200", "registry_path": "docs/guides/resources--cloud_credentials--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:cloud_credentials:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["aws assume role"], "anchor": "section", "description": "AWS Assume Role to Handle Delegated Access.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "conflicts"}, {"anchor": "schema-aws_assume_role--custom_external_id", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_optional", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:custom_external_id,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_assume_role:ConflictingObjectAttributes:external_id_is_optional,external_id_is_tenant_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role:external_id_is_tenant_id", "type": "conflicts"}, {"anchor": "schema-aws_assume_role--duration_seconds", "enforcement": "provider-schema", "group": "aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_assume_role--role_arn", "enforcement": "provider-schema", "group": "aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "requires"}, {"anchor": "schema-aws_assume_role--session_name", "enforcement": "provider-schema", "group": "aws_assume_role:RequiredObjectAttributes:duration_seconds,role_arn,session_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_assume_role", "type": "requires"}], "schema_path": ["aws_assume_role"], "syntax": "block", "type": "object"}, {"aliases": ["aws secret key"], "anchor": "section", "description": "AWS Programmatic Access Credentials type.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_secret_key--access_key", "enforcement": "provider-schema", "group": "aws_secret_key:RequiredObjectAttributes:access_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:aws_secret_key", "type": "requires"}], "schema_path": ["aws_secret_key"], "syntax": "block", "type": "object"}, {"aliases": ["azure client secret"], "anchor": "section", "description": "Azure Credentials Client Secret type.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_client_secret--client_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}, {"anchor": "schema-azure_client_secret--subscription_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}, {"anchor": "schema-azure_client_secret--tenant_id", "enforcement": "provider-schema", "group": "azure_client_secret:RequiredObjectAttributes:client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_client_secret", "type": "requires"}], "schema_path": ["azure_client_secret"], "syntax": "block", "type": "object"}, {"aliases": ["azure pfx certificate"], "anchor": "section", "description": "Azure Credentials Client Certificate type.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-azure_pfx_certificate--certificate_url", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--client_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--subscription_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}, {"anchor": "schema-azure_pfx_certificate--tenant_id", "enforcement": "provider-schema", "group": "azure_pfx_certificate:RequiredObjectAttributes:certificate_url,client_id,subscription_id,tenant_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_credentials:properties:azure_pfx_certificate", "type": "requires"}], "schema_path": ["azure_pfx_certificate"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:cloud_credentials:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:cloud_credentials:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["gcp cred file"], "anchor": "section", "description": "GCP Credentials type.", "document_id": "xcsh-docs:resources:cloud_credentials:properties:gcp_cred_file", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gcp_cred_file"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:cloud_credentials:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:cloud_credentials:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:cloud_credentials:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:cloud_credentials:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:cloud_credentials:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_cloud_credentials.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/): complete subsection reference.

- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/): complete subsection reference.

- [azure_client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/): complete subsection reference.

- [azure_pfx_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Cloud Credentials. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Cloud Credentials is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/#schema-annotations) |
| `aws_assume_role` | [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/#section) |
| `aws_assume_role.custom_external_id` | [aws_assume_role.custom_external_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/#schema-aws_assume_role--custom_external_id) |
| `aws_assume_role.duration_seconds` | [aws_assume_role.duration_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/#schema-aws_assume_role--duration_seconds) |
| `aws_assume_role.external_id_is_optional` | [aws_assume_role.external_id_is_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/external_id_is_optional/#section) |
| `aws_assume_role.external_id_is_tenant_id` | [aws_assume_role.external_id_is_tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/external_id_is_tenant_id/#section) |
| `aws_assume_role.role_arn` | [aws_assume_role.role_arn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/#schema-aws_assume_role--role_arn) |
| `aws_assume_role.session_name` | [aws_assume_role.session_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/#schema-aws_assume_role--session_name) |
| `aws_assume_role.session_tags` | [aws_assume_role.session_tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/#schema-aws_assume_role--session_tags) |
| `aws_secret_key` | [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/#section) |
| `aws_secret_key.access_key` | [aws_secret_key.access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/#schema-aws_secret_key--access_key) |
| `aws_secret_key.secret_key` | [aws_secret_key.secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/#section) |
| `aws_secret_key.secret_key.blindfold_secret_info` | [aws_secret_key.secret_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/#section) |
| `aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/#schema-aws_secret_key--secret_key--blindfold_secret_info--decryption_provider) |
| `aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_secret_key.secret_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/#schema-aws_secret_key--secret_key--blindfold_secret_info--location) |
| `aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_secret_key.secret_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/blindfold_secret_info/#schema-aws_secret_key--secret_key--blindfold_secret_info--store_provider) |
| `aws_secret_key.secret_key.clear_secret_info` | [aws_secret_key.secret_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/clear_secret_info/#section) |
| `aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_secret_key.secret_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/clear_secret_info/#schema-aws_secret_key--secret_key--clear_secret_info--provider_ref) |
| `aws_secret_key.secret_key.clear_secret_info.url` | [aws_secret_key.secret_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/secret_key/clear_secret_info/#schema-aws_secret_key--secret_key--clear_secret_info--url) |
| `azure_client_secret` | [azure_client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/#section) |
| `azure_client_secret.client_id` | [azure_client_secret.client_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/#schema-azure_client_secret--client_id) |
| `azure_client_secret.client_secret` | [azure_client_secret.client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/#section) |
| `azure_client_secret.client_secret.blindfold_secret_info` | [azure_client_secret.client_secret.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/blindfold_secret_info/#section) |
| `azure_client_secret.client_secret.blindfold_secret_info.decryption_provider` | [azure_client_secret.client_secret.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/blindfold_secret_info/#schema-azure_client_secret--client_secret--blindfold_secret_info--decryption_provider) |
| `azure_client_secret.client_secret.blindfold_secret_info.location` | [azure_client_secret.client_secret.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/blindfold_secret_info/#schema-azure_client_secret--client_secret--blindfold_secret_info--location) |
| `azure_client_secret.client_secret.blindfold_secret_info.store_provider` | [azure_client_secret.client_secret.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/blindfold_secret_info/#schema-azure_client_secret--client_secret--blindfold_secret_info--store_provider) |
| `azure_client_secret.client_secret.clear_secret_info` | [azure_client_secret.client_secret.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/clear_secret_info/#section) |
| `azure_client_secret.client_secret.clear_secret_info.provider_ref` | [azure_client_secret.client_secret.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/clear_secret_info/#schema-azure_client_secret--client_secret--clear_secret_info--provider_ref) |
| `azure_client_secret.client_secret.clear_secret_info.url` | [azure_client_secret.client_secret.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/client_secret/clear_secret_info/#schema-azure_client_secret--client_secret--clear_secret_info--url) |
| `azure_client_secret.subscription_id` | [azure_client_secret.subscription_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/#schema-azure_client_secret--subscription_id) |
| `azure_client_secret.tenant_id` | [azure_client_secret.tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/#schema-azure_client_secret--tenant_id) |
| `azure_pfx_certificate` | [azure_pfx_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/#section) |
| `azure_pfx_certificate.certificate_url` | [azure_pfx_certificate.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/#schema-azure_pfx_certificate--certificate_url) |
| `azure_pfx_certificate.client_id` | [azure_pfx_certificate.client_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/#schema-azure_pfx_certificate--client_id) |
| `azure_pfx_certificate.password` | [azure_pfx_certificate.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/#section) |
| `azure_pfx_certificate.password.blindfold_secret_info` | [azure_pfx_certificate.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/blindfold_secret_info/#section) |
| `azure_pfx_certificate.password.blindfold_secret_info.decryption_provider` | [azure_pfx_certificate.password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/blindfold_secret_info/#schema-azure_pfx_certificate--password--blindfold_secret_info--decryption_provider) |
| `azure_pfx_certificate.password.blindfold_secret_info.location` | [azure_pfx_certificate.password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/blindfold_secret_info/#schema-azure_pfx_certificate--password--blindfold_secret_info--location) |
| `azure_pfx_certificate.password.blindfold_secret_info.store_provider` | [azure_pfx_certificate.password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/blindfold_secret_info/#schema-azure_pfx_certificate--password--blindfold_secret_info--store_provider) |
| `azure_pfx_certificate.password.clear_secret_info` | [azure_pfx_certificate.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/clear_secret_info/#section) |
| `azure_pfx_certificate.password.clear_secret_info.provider_ref` | [azure_pfx_certificate.password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/clear_secret_info/#schema-azure_pfx_certificate--password--clear_secret_info--provider_ref) |
| `azure_pfx_certificate.password.clear_secret_info.url` | [azure_pfx_certificate.password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/password/clear_secret_info/#schema-azure_pfx_certificate--password--clear_secret_info--url) |
| `azure_pfx_certificate.subscription_id` | [azure_pfx_certificate.subscription_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/#schema-azure_pfx_certificate--subscription_id) |
| `azure_pfx_certificate.tenant_id` | [azure_pfx_certificate.tenant_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/#schema-azure_pfx_certificate--tenant_id) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/#schema-disable) |
| `gcp_cred_file` | [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/#section) |
| `gcp_cred_file.credential_file` | [gcp_cred_file.credential_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/#section) |
| `gcp_cred_file.credential_file.blindfold_secret_info` | [gcp_cred_file.credential_file.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/#section) |
| `gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/#schema-gcp_cred_file--credential_file--blindfold_secret_info--decryption_provider) |
| `gcp_cred_file.credential_file.blindfold_secret_info.location` | [gcp_cred_file.credential_file.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/#schema-gcp_cred_file--credential_file--blindfold_secret_info--location) |
| `gcp_cred_file.credential_file.blindfold_secret_info.store_provider` | [gcp_cred_file.credential_file.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/blindfold_secret_info/#schema-gcp_cred_file--credential_file--blindfold_secret_info--store_provider) |
| `gcp_cred_file.credential_file.clear_secret_info` | [gcp_cred_file.credential_file.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/clear_secret_info/#section) |
| `gcp_cred_file.credential_file.clear_secret_info.provider_ref` | [gcp_cred_file.credential_file.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/clear_secret_info/#schema-gcp_cred_file--credential_file--clear_secret_info--provider_ref) |
| `gcp_cred_file.credential_file.clear_secret_info.url` | [gcp_cred_file.credential_file.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/credential_file/clear_secret_info/#schema-gcp_cred_file--credential_file--clear_secret_info--url) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/#schema-namespace) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/timeouts/#schema-timeouts--update) |

## Next pages

- [aws_assume_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_assume_role/)
- [aws_secret_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/aws_secret_key/)
- [azure_client_secret](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_client_secret/)
- [azure_pfx_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/azure_pfx_certificate/)
- [gcp_cred_file](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/gcp_cred_file/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/timeouts/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
