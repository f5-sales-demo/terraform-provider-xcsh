---
page_title: "vsite_refs"
subcategory: "Container"
description: "Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this virtual-site is used select sites on which to instantiate the Kubernetes API resource object."
xcsh_docs: {"aliases": ["vsite refs"], "body_bytes": 5225, "body_sha256": "sha256:40a82ffecb2badda8219db38eae243030e5cf8f596d5a10c8939c6041136eb52", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs", "parent_id": "xcsh-docs:data-sources:virtual_k8s:reference", "path": "documentation/data-sources/virtual_k8s/properties/vsite_refs/index.md", "product": "distributed-cloud", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2220313111103300-3321111021211031-3032231310121133-3130001012030101-1022230330211132-0302212212212220-3013321022121013-3003321003103022", "registry_path": "docs/guides/data-sources--virtual_k8s--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vsite_refs"], "schema_version": 1, "sections": [{"aliases": ["vsite refs kind"], "anchor": "schema-vsite_refs--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "kind", "scope_path": ["vsite_refs"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["vsite_refs", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["vsite refs name"], "anchor": "schema-vsite_refs--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "name", "scope_path": ["vsite_refs"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["vsite_refs", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["vsite refs namespace"], "anchor": "schema-vsite_refs--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "namespace", "scope_path": ["vsite_refs"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["vsite_refs", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["vsite refs tenant"], "anchor": "schema-vsite_refs--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "tenant", "scope_path": ["vsite_refs"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["vsite_refs", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["vsite refs uid"], "anchor": "schema-vsite_refs--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:data-sources:virtual_k8s:properties:vsite_refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "reference_identity": {"member": "uid", "scope_path": ["vsite_refs"], "source": "receipt-pinned-schema-identity", "upstream_message": "ves.io.schema.ObjectRefType", "version": 1}, "relationships": [], "schema_path": ["vsite_refs", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/properties/vsite_refs/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this virtual-site is used select sites on which to instantiate the Kubernetes API resource object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vsite_refs

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_k8s/properties/)
- vsite_refs

<a id="section"></a>

Type: `"list"`. Computed.

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

## Direct properties

<a id="schema-vsite_refs--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-vsite_refs--name"></a>

### name property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-vsite_refs--namespace"></a>

### namespace property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-vsite_refs--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="schema-vsite_refs--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
