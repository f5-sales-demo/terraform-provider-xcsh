---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 3579, "body_sha256": "sha256:5d93188d0b672921f039244c8ebcd53752bc1ce6ea22b62ed67cc939b423d293", "canonical_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:port", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:port:info"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:port", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports", "path": "docs/guides/data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--port.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [stateful_service](data-sources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](data-sources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports.md)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="section"></a>

Type: `"single"`. Computed.

Port. Port of the workload.

Upstream description:

Port of the workload.

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

- [info](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--port--info.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--multi_ports--ports--port--name"></a>

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

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports--port--info.md)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--properties--stateful_service--advertise_options--advertise_on_public--multi_ports--ports.md)
- [xcsh_workload](../data-sources/workload.md)
