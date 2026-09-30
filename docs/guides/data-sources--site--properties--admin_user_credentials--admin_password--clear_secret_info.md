---
page_title: "admin_user_credentials.admin_password.clear_secret_info"
subcategory: "Infrastructure"
description: "admin_user_credentials.admin_password.clear_secret_info for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1445, "body_sha256": "sha256:92b3107d4b7e0d47f812f3c3583b526cf7d7cce3bc63a8a219168e900d49a4c8", "canonical_id": "xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password:clear_secret_info", "parent_id": "xcsh-docs:data-sources:site:properties:admin_user_credentials:admin_password", "path": "docs/guides/data-sources--site--properties--admin_user_credentials--admin_password--clear_secret_info.md", "provider_name": "site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["admin_user_credentials", "admin_password", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/properties/admin_user_credentials/admin_password/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "admin_user_credentials.admin_password.clear_secret_info for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# admin_user_credentials.admin_password.clear_secret_info

Breadcrumbs:

- [xcsh_site](../data-sources/site.md)
- [Property reference](data-sources--site--reference.md)
- [admin_user_credentials](data-sources--site--properties--admin_user_credentials.md)
- [admin_user_credentials.admin_password](data-sources--site--properties--admin_user_credentials--admin_password.md)
- admin_user_credentials.admin_password.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

## Direct properties

<a id="schema-admin_user_credentials--admin_password--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-admin_user_credentials--admin_password--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

## Next pages

- [admin_user_credentials.admin_password](data-sources--site--properties--admin_user_credentials--admin_password.md)
- [xcsh_site](../data-sources/site.md)
