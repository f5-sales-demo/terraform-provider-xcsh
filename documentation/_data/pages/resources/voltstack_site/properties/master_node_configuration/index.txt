---
page_title: "master_node_configuration"
subcategory: ""
description: "Configuration of master nodes."
xcsh_docs: {"aliases": ["master node configuration"], "body_bytes": 4514, "body_sha256": "sha256:a8afa2e5f201f8f9e4bbaea3927dac7bab64aee5508610d9ceed6584756ee639", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:master_node_configuration", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/master_node_configuration/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2231013002202321-1310031001031301-2022032301131330-1102123002203131-0101322113031031-0102323000220120-2121200313301302-1032303230301033", "registry_path": "docs/guides/resources--voltstack_site--reference--group-010.md", "relationships": [{"anchor": "schema-master_node_configuration--name", "enforcement": "provider-schema", "group": "master_node_configuration:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:master_node_configuration", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["master_node_configuration"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-master_node_configuration--name", "description": "Names of master node.", "document_id": "xcsh-docs:resources:voltstack_site:properties:master_node_configuration", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["master_node_configuration", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["public ip"], "anchor": "schema-master_node_configuration--public_ip", "description": "IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.", "document_id": "xcsh-docs:resources:voltstack_site:properties:master_node_configuration", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["master_node_configuration", "public_ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/master_node_configuration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration of master nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# master_node_configuration

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- master_node_configuration

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3"
  }
}
```

Terraform syntax:

```terraform
master_node_configuration {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-master_node_configuration--name"></a>

### name property

Type: `"string"`. Optional.

Name. Names of master node.

Upstream description:

Names of master node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="schema-master_node_configuration--public_ip"></a>

### public_ip property

Type: `"string"`. Optional.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
