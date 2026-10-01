---
page_title: "stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 4362, "body_sha256": "sha256:e909c8e616dbc04ff57ff51ff82ea1581206324bae995638f74ffc6d76ff3ce5", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports:info"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports:ports", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_in_cluster:multi_ports", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/multi_ports/ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster.md)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--multi_ports.md)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

## Direct properties

- [info](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--name"></a>

### name property

Type: `"string"`. Optional.

Name. Name of the Port.

Upstream description:

Name of the Port.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

## Next pages

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--multi_ports--ports--info.md)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--properties--stateful_service--advertise_options--advertise_in_cluster--multi_ports.md)
- [xcsh_workload](../resources/workload.md)
