---
page_title: "voltstack_cluster.site_local_network.existing_network"
subcategory: "Infrastructure"
description: "voltstack_cluster.site_local_network.existing_network for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2778, "body_sha256": "sha256:4525e0a7d66346942a157ef90f711a2dfa18d2c9067ef70d9fbadd5eff80db54", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "child_ids": [], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "path": "docs/guides/data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network--existing_network.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "site_local_network", "existing_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/existing_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.site_local_network.existing_network for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.site_local_network.existing_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md)
- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network.md)
- voltstack_cluster.site_local_network.existing_network

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

## Direct properties

<a id="schema-voltstack_cluster--site_local_network--existing_network--name"></a>

### name property

Type: `"string"`. Computed.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
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
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--properties--voltstack_cluster--site_local_network.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
