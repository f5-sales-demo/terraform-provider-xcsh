---
page_title: "api_protection_rules.api_groups_rules"
subcategory: "Load Balancing"
description: "api_protection_rules.api_groups_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6724, "body_sha256": "sha256:ea741173918490c9214c34234b08699f56fc49c885f36fbf12a27d449af14042", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:client_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:request_matcher"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_protection_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules.api_groups_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_protection_rules](data-sources--http_loadbalancer--properties--api_protection_rules.md)
- api_protection_rules.api_groups_rules

<a id="section"></a>

Type: `"list"`. Computed.

Category includes rules per API group or Server URL. For API groups, refer to API Definition which
includes API groups derived from uploaded swaggers.

Upstream description:

This category includes rules per API group or Server URL. For API groups, refer to API Definition
which includes API groups derived from uploaded swaggers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

## Direct properties

- [action](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action.md): complete subsection reference.

- [any_domain](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--any_domain.md): complete subsection reference.

<a id="schema-api_protection_rules--api_groups_rules--api_group"></a>

### api_group property

Type: `"string"`. Computed.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-api_protection_rules--api_groups_rules--base_path"></a>

### base_path property

Type: `"string"`. Computed.

Base Path. Prefix of the request path. For example: /v1.

Upstream description:

Prefix of the request path. For example: /v1.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--client_matcher.md): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--metadata.md): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher.md): complete subsection reference.

<a id="schema-api_protection_rules--api_groups_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

## Next pages

- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--action.md)
- [api_protection_rules.api_groups_rules.any_domain](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--any_domain.md)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--client_matcher.md)
- [api_protection_rules.api_groups_rules.metadata](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--metadata.md)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--properties--api_protection_rules--api_groups_rules--request_matcher.md)
- [api_protection_rules](data-sources--http_loadbalancer--properties--api_protection_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
