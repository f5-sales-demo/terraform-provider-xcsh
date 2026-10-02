---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_managed_client_customer_support_comments."
xcsh_docs: {"aliases": ["managed client customer support comments"], "body_bytes": 5069, "body_sha256": "sha256:d429176a8fe8b44c37e4e0753641d6dcdfee5bcca4dce07b8300f9fc290cb842", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:managed_client_customer_support_comments:reference", "parent_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:fundamentals", "path": "documentation/data-sources/managed_client_customer_support_comments/properties/index.md", "product": "distributed-cloud", "provider_name": "managed_client_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0000022210111311-2312013223102231-1230223030030013-3002331301021100-0113020132311001-2203322121231303-1313033333102303-3220002111232101", "registry_path": "docs/guides/data-sources--managed_client_customer_support_comments--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["all comments returned"], "anchor": "schema-all_comments_returned", "description": "Indicates if all comments for the issue have been returned.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["all_comments_returned"], "syntax": "attribute", "type": "bool"}, {"aliases": ["comments"], "anchor": "section", "description": "List of comments on the customer support ticket.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:properties:comments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["comments"], "syntax": "attribute", "type": "object"}, {"aliases": ["created until timestamp"], "anchor": "schema-created_until_timestamp", "description": "Filter to retrieve comments created up to the specified timestamp.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["created_until_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["tp id"], "anchor": "schema-tp_id", "description": "Tpid ID assigned to this ticket by Third Party.", "document_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tp_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/managed_client_customer_support_comments/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_managed_client_customer_support_comments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/)
- Property reference

## Direct properties

<a id="schema-all_comments_returned"></a>

### all_comments_returned property

Type: `"bool"`. Computed.

Indicates if all comments for the issue have been returned.

- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/): complete subsection reference.

<a id="schema-created_until_timestamp"></a>

### created_until_timestamp property

Type: `"string"`. Optional.

Filter to retrieve comments created up to the specified timestamp.

<a id="schema-tp_id"></a>

### tp_id property

Type: `"string"`. Required.

Tpid ID assigned to this ticket by Third Party.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `all_comments_returned` | [all_comments_returned](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/#schema-all_comments_returned) |
| `comments` | [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#section) |
| `comments.attachment_ids` | [comments.attachment_ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#schema-comments--attachment_ids) |
| `comments.attachments_info` | [comments.attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/#section) |
| `comments.attachments_info.attachment` | [comments.attachments_info.attachment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--attachment) |
| `comments.attachments_info.content_type` | [comments.attachments_info.content_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--content_type) |
| `comments.attachments_info.filename` | [comments.attachments_info.filename](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--filename) |
| `comments.attachments_info.tp_id` | [comments.attachments_info.tp_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/attachments_info/#schema-comments--attachments_info--tp_id) |
| `comments.author_email` | [comments.author_email](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#schema-comments--author_email) |
| `comments.author_name` | [comments.author_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#schema-comments--author_name) |
| `comments.comment_id` | [comments.comment_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#schema-comments--comment_id) |
| `comments.created_at` | [comments.created_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#schema-comments--created_at) |
| `comments.html` | [comments.html](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#schema-comments--html) |
| `comments.plain_text` | [comments.plain_text](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/#schema-comments--plain_text) |
| `created_until_timestamp` | [created_until_timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/#schema-created_until_timestamp) |
| `tp_id` | [tp_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/#schema-tp_id) |

## Next pages

- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/comments/)
- [xcsh_managed_client_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/)
