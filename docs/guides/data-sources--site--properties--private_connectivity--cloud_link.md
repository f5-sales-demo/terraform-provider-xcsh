---
page_title: "private_connectivity.cloud_link"
subcategory: "Infrastructure"
description: "private_connectivity.cloud_link for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1320, "body_sha256": "sha256:a0c97a252defef59923c9dddc67ae34648ac0372c218086f54f5a0b949d9b60f", "canonical_id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:private_connectivity:cloud_link", "parent_id": "xcsh-docs:data-sources:site:properties:private_connectivity", "path": "docs/guides/data-sources--site--properties--private_connectivity--cloud_link.md", "provider_name": "site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["private_connectivity", "cloud_link"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/private_connectivity/cloud_link/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connectivity.cloud_link for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity.cloud_link

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- [private_connectivity](data-sources--site--properties--private_connectivity.md)
- private_connectivity.cloud_link

<a id="section"></a>

Type: `"single"`. Computed.

Information related to cloud link used by the site.

## Direct properties

<a id="schema-private_connectivity--cloud_link--name"></a>

### name property

Type: `"string"`. Computed.

Name of the the CloudLink used with this site.

<a id="schema-private_connectivity--cloud_link--state"></a>

### state property

Type: `"string"`. Computed.

\[Enum: UP|DOWN|DEGRADED|NOT\_APPLICABLE\] State of the CloudLink connections - UP: Up CloudLink and
their corresponding Direct Connect connections are up and healthy - DOWN: Down CloudLink and their
corresponding Direct Connect connections are down - DEGRADED: Degraded Some of Direct Connect
connections with the CloudLink are down .. Possible values are \`UP\`, \`DOWN\`, \`DEGRADED\`,
\`NOT\_APPLICABLE\`. Defaults to \`UP\`.

## Next pages

- [private_connectivity](data-sources--site--properties--private_connectivity.md)
- [xcsh_site](../data-sources/site.md)
