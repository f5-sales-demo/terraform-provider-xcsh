---
page_title: "service.advertise_options.advertise_in_cluster.multi_ports.ports"
subcategory: "Container"
description: "service.advertise_options.advertise_in_cluster.multi_ports.ports for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3756, "body_sha256": "sha256:f62a4e183cf076c3a213434eabd761ae67ca515fe86133657c008ded98372872", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports:info"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports:ports", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_in_cluster:multi_ports", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports--ports.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_in_cluster", "multi_ports", "ports"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_in_cluster/multi_ports/ports/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_in_cluster.multi_ports.ports for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_in_cluster.multi_ports.ports

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_in_cluster](data-sources--workload--properties--service--advertise_options--advertise_in_cluster.md)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports.md)
- service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="section"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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

## Direct properties

- [info](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports--ports--info.md): complete subsection reference.

<a id="schema-service--advertise_options--advertise_in_cluster--multi_ports--ports--name"></a>

### name property

Type: `"string"`. Computed.

Name. Name of the Port.

Upstream description:

Name of the Port.

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

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports--ports--info.md)
- [service.advertise_options.advertise_in_cluster.multi_ports](data-sources--workload--properties--service--advertise_options--advertise_in_cluster--multi_ports.md)
- [xcsh_workload](../data-sources/workload.md)
