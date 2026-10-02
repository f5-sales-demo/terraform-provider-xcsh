---
page_title: "simple_service.container.readiness_check.http_health_check.port"
subcategory: "Container"
description: "Port"
xcsh_docs: {"aliases": ["simple service container readiness check http health check port"], "body_bytes": 4177, "body_sha256": "sha256:07f09a6d156ab016fd086fbeb14341fc477d9084757677575bc13f2c4e4b9d85", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:http_health_check:port", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:http_health_check", "path": "documentation/data-sources/workload/properties/simple_service/container/readiness_check/http_health_check/port/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3012203012020033-3212012332003203-1011303232113213-1311031123301011-1333223031113102-0010112133103022-1232333111202001-3023211311112032", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "readiness_check", "http_health_check", "port"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-simple_service--container--readiness_check--http_health_check--port--name", "description": "Exclusive with Port Name.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:http_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "http_health_check", "port", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["num"], "anchor": "schema-simple_service--container--readiness_check--http_health_check--port--num", "description": "Exclusive with Port number.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:http_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "http_health_check", "port", "num"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/container/readiness_check/http_health_check/port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Port", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.readiness_check.http_health_check.port

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/)
- [simple_service.container.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/)
- [simple_service.container.readiness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/http_health_check/)
- simple_service.container.readiness_check.http_health_check.port

<a id="section"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

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

## Direct properties

<a id="schema-simple_service--container--readiness_check--http_health_check--port--name"></a>

### name property

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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

<a id="schema-simple_service--container--readiness_check--http_health_check--port--num"></a>

### num property

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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

- [simple_service.container.readiness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/http_health_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
