---
page_title: "infra.hw_info.cpu"
subcategory: ""
description: "CPU information."
xcsh_docs: {"aliases": ["infra hw info cpu"], "body_bytes": 4115, "body_sha256": "sha256:c62e2eeb021bcd4e02d470fba6d5ae55ebdc7fa2fe4ea10620842de1c605bb9a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "documentation/data-sources/registration/properties/infra/hw_info/cpu/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2310202130323300-2033301303131032-0201000311132033-1222111330021332-0312020321111232-1013202023021311-2030212020012223-3133012100322023", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "cpu"], "schema_version": 1, "sections": [{"aliases": ["cache"], "anchor": "schema-infra--hw_info--cpu--cache", "description": "CPU cache size in KB.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "cpu", "cache"], "syntax": "attribute", "type": "number"}, {"aliases": ["cores"], "anchor": "schema-infra--hw_info--cpu--cores", "description": "Number of physical CPU cores.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "cpu", "cores"], "syntax": "attribute", "type": "number"}, {"aliases": ["cpus"], "anchor": "schema-infra--hw_info--cpu--cpus", "description": "Number of physical CPUs.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "cpu", "cpus"], "syntax": "attribute", "type": "number"}, {"aliases": ["model"], "anchor": "schema-infra--hw_info--cpu--model", "description": "CPU model", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "cpu", "model"], "syntax": "attribute", "type": "string"}, {"aliases": ["speed"], "anchor": "schema-infra--hw_info--cpu--speed", "description": "CPU clock rate in MHz.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "cpu", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["threads"], "anchor": "schema-infra--hw_info--cpu--threads", "description": "Number of logical (HT) CPU cores.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "cpu", "threads"], "syntax": "attribute", "type": "number"}, {"aliases": ["vendor"], "anchor": "schema-infra--hw_info--cpu--vendor", "description": "CPU vendor.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "cpu", "vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/cpu/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CPU information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.cpu

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- infra.hw_info.cpu

<a id="section"></a>

Type: `"single"`. Computed.

CPU Information. CPU information.

Upstream description:

CPU information.

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

<a id="schema-infra--hw_info--cpu--cache"></a>

### cache property

Type: `"number"`. Computed.

Cache. CPU cache size in KB.

Upstream description:

CPU cache size in KB.

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

<a id="schema-infra--hw_info--cpu--cores"></a>

### cores property

Type: `"number"`. Computed.

Cores. Number of physical CPU cores.

Upstream description:

Number of physical CPU cores.

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

<a id="schema-infra--hw_info--cpu--cpus"></a>

### cpus property

Type: `"number"`. Computed.

CPUs. Number of physical CPUs.

Upstream description:

Number of physical CPUs.

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

<a id="schema-infra--hw_info--cpu--model"></a>

### model property

Type: `"string"`. Computed.

Model. CPU model

Upstream description:

CPU model

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-infra--hw_info--cpu--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. CPU clock rate in MHz.

Upstream description:

CPU clock rate in MHz.

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

<a id="schema-infra--hw_info--cpu--threads"></a>

### threads property

Type: `"number"`. Computed.

Threads. Number of logical (HT) CPU cores.

Upstream description:

Number of logical (HT) CPU cores.

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

<a id="schema-infra--hw_info--cpu--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. CPU vendor.

Upstream description:

CPU vendor.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
