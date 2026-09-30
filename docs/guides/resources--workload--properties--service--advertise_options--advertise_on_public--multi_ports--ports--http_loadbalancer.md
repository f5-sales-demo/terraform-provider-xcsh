---
page_title: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 7222, "body_sha256": "sha256:75493745a0816f9e6ae54fe0297708bafc3b10e124a1358151c5bc6b7865e58d", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:default_route", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:http", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:https_auto_cert", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:multi_ports:ports", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports.md)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_route](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--default_route.md): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--domains"></a>

### domains property

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http.md): complete subsection reference.

- [https](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https.md): complete subsection reference.

- [https_auto_cert](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert.md): complete subsection reference.

- [specific_routes](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--default_route.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--http.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--https_auto_cert.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes.md)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--properties--service--advertise_options--advertise_on_public--multi_ports--ports.md)
- [xcsh_workload](../resources/workload.md)
