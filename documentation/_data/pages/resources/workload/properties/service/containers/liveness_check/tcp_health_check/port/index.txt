---
page_title: "service.containers.liveness_check.tcp_health_check.port"
subcategory: "Container"
description: "Port"
xcsh_docs: {"aliases": ["service containers liveness check tcp health check port"], "body_bytes": 4688, "body_sha256": "sha256:72f94e68e957e2dc90236b1aef05c8f0e859f5da4f2540112652d1ab1510e22a", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check:port", "parent_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check", "path": "documentation/resources/workload/properties/service/containers/liveness_check/tcp_health_check/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0033013130231022-2310211300202230-3301310301233111-0120011223101333-0032303121201113-3333330023221310-1033121202312312-2323310020123331", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "schema-service--containers--liveness_check--tcp_health_check--port--name", "enforcement": "provider-schema", "group": "service.containers.liveness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check:port", "type": "conflicts"}, {"anchor": "schema-service--containers--liveness_check--tcp_health_check--port--num", "enforcement": "provider-schema", "group": "service.containers.liveness_check.tcp_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check:port", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "containers", "liveness_check", "tcp_health_check", "port"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-service--containers--liveness_check--tcp_health_check--port--name", "description": "Exclusive with Port Name.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check:port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "tcp_health_check", "port", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["num"], "anchor": "schema-service--containers--liveness_check--tcp_health_check--port--num", "description": "Exclusive with Port number.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:tcp_health_check:port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "tcp_health_check", "port", "num"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/containers/liveness_check/tcp_health_check/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Port", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers.liveness_check.tcp_health_check.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/)
- [service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/)
- [service.containers.liveness_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/tcp_health_check/)
- service.containers.liveness_check.tcp_health_check.port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--containers--liveness_check--tcp_health_check--port--name"></a>

### name property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="schema-service--containers--liveness_check--tcp_health_check--port--num"></a>

### num property

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [service.containers.liveness_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/tcp_health_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
