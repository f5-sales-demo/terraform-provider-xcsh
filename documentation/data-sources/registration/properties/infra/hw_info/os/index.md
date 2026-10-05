---
page_title: "infra.hw_info.os"
subcategory: ""
description: "Details of Operating System."
xcsh_docs: {"aliases": ["infra hw info os"], "body_bytes": 4694, "body_sha256": "sha256:65602542664534645019661bc0c2d8c13edb32fb77b4cf58def0cd3895c64839", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "parent_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "path": "documentation/data-sources/registration/properties/infra/hw_info/os/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2023302323122322-3003210101112121-2100312221020002-2332203332311233-0010331003002231-0120201232012120-2332013030003023-0101232230013331", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "os"], "schema_version": 1, "sections": [{"aliases": ["infra hw info os architecture"], "anchor": "schema-infra--hw_info--os--architecture", "description": "Architecture of OS.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "os", "architecture"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info os name"], "anchor": "schema-infra--hw_info--os--name", "description": "Name of OS.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "os", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info os release"], "anchor": "schema-infra--hw_info--os--release", "description": "Release of the OS.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "os", "release"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info os vendor"], "anchor": "schema-infra--hw_info--os--vendor", "description": "Vendor of OS.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "os", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra hw info os version"], "anchor": "schema-infra--hw_info--os--version", "description": "Version of OS.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "os", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/hw_info/os/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Details of Operating System.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.os

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- infra.hw_info.os

<a id="section"></a>

Type: `"single"`. Computed.

OS. Details of Operating System.

Upstream description:

Details of Operating System.

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

<a id="schema-infra--hw_info--os--architecture"></a>

### architecture property

Type: `"string"`. Computed.

Architecture. Architecture of OS.

Upstream description:

Architecture of OS.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-infra--hw_info--os--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of OS.

Upstream description:

Name of OS.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--os--release"></a>

### release property

Type: `"string"`. Computed.

Release. Release of the OS.

Upstream description:

Release of the OS.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-infra--hw_info--os--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. Vendor of OS.

Upstream description:

Vendor of OS.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-infra--hw_info--os--version"></a>

### version property

Type: `"string"`. Computed.

Version. Version of OS.

Upstream description:

Version of OS.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
