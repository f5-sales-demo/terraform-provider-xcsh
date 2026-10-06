---
page_title: "Property reference"
subcategory: "Container"
description: "Property reference for xcsh_virtual_k8s."
xcsh_docs: {"aliases": ["virtual k8s"], "body_bytes": 10045, "body_sha256": "sha256:e539a38d259d4149bbcf8435b23cdaa869619f5e27a2da4beb46c8f14489d78e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:virtual_k8s:properties:default_flavor_ref", "xcsh-docs:data-sources:virtual_k8s:properties:disabled", "xcsh-docs:data-sources:virtual_k8s:properties:isolated", "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_k8s:reference", "parent_id": "xcsh-docs:data-sources:virtual_k8s:fundamentals", "path": "documentation/data-sources/virtual_k8s/properties/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130", "registry_path": "docs/guides/data-sources--virtual_k8s--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:virtual_k8s:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["default flavor ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:default_flavor_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_flavor_ref"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:virtual_k8s:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:virtual_k8s:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["isolated"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:isolated", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["isolated"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:virtual_k8s:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:virtual_k8s:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:virtual_k8s:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["vsite refs"], "anchor": "section", "description": "Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this virtual-site is used select sites on which to instantiate the Kubernetes API resource object.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["vsite_refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_virtual_k8s.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

- [default_flavor_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/default_flavor_ref/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the VirtualK8S.

Additional upstream details:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/disabled/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/isolated/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

Name of the VirtualK8S.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Namespace where the VirtualK8S exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [vsite_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/vsite_refs/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/#schema-annotations) |
| `default_flavor_ref` | [default_flavor_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/default_flavor_ref/#section) |
| `default_flavor_ref.name` | [default_flavor_ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/default_flavor_ref/#schema-default_flavor_ref--name) |
| `default_flavor_ref.namespace` | [default_flavor_ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/default_flavor_ref/#schema-default_flavor_ref--namespace) |
| `default_flavor_ref.tenant` | [default_flavor_ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/default_flavor_ref/#schema-default_flavor_ref--tenant) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/#schema-description) |
| `disabled` | [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/disabled/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/#schema-id) |
| `isolated` | [isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/isolated/#section) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/#schema-namespace) |
| `vsite_refs` | [vsite_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/vsite_refs/#section) |
| `vsite_refs.kind` | [vsite_refs.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/vsite_refs/#schema-vsite_refs--kind) |
| `vsite_refs.name` | [vsite_refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/vsite_refs/#schema-vsite_refs--name) |
| `vsite_refs.namespace` | [vsite_refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/vsite_refs/#schema-vsite_refs--namespace) |
| `vsite_refs.tenant` | [vsite_refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/vsite_refs/#schema-vsite_refs--tenant) |
| `vsite_refs.uid` | [vsite_refs.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/vsite_refs/#schema-vsite_refs--uid) |
