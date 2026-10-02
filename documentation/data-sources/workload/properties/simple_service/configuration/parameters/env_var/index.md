---
page_title: "simple_service.configuration.parameters.env_var"
subcategory: "Container"
description: "Environment Variable."
xcsh_docs: {"aliases": ["simple service configuration parameters env var"], "body_bytes": 3762, "body_sha256": "sha256:f6201eb5053c80721f98da73ea38c4e96bcc71f339bece8109c4e3c755ded601", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:env_var", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters", "path": "documentation/data-sources/workload/properties/simple_service/configuration/parameters/env_var/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0311103231201330-2203202132201132-2231032300112002-0131121232112210-3333000331023313-2022022013000310-0031332233300200-3121130302021303", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "configuration", "parameters", "env_var"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-simple_service--configuration--parameters--env_var--name", "description": "Name of Environment Variable.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:env_var", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "env_var", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "schema-simple_service--configuration--parameters--env_var--value", "description": "Value of Environment Variable.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:configuration:parameters:env_var", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "configuration", "parameters", "env_var", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/configuration/parameters/env_var/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Environment Variable.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.configuration.parameters.env_var

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/)
- [simple_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/)
- simple_service.configuration.parameters.env_var

<a id="section"></a>

Type: `"single"`. Computed.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

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

## Direct properties

<a id="schema-simple_service--configuration--parameters--env_var--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-simple_service--configuration--parameters--env_var--value"></a>

### value property

Type: `"string"`. Computed.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

## Next pages

- [simple_service.configuration.parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/configuration/parameters/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
