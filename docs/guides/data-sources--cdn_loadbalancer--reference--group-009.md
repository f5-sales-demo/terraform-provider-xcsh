---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-de29098d2fdb5ea8e1a91c2ae9b133c1ad8f7d4ed34e2e79dccc3a8379b1e726"></a>

## client_side_defense — client_side_defense / ba208314a041 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- client_side_defense

<a id="canonical-c3f925fa8d15f95853ebddc4767aff82037e6a0be261a280b5308136b5b64182"></a>

Type: `"single"`. Computed.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

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

OneOf alternatives in this subsection:

- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-c3f925fa8d15f95853ebddc4767aff82037e6a0be261a280b5308136b5b64182)
- [disable_client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a885b2a10be231e4132d2d07c964ee4f3a3b106473d40194f7fbfd22fbd77c04)

Select alternatives according to the provider validators above.

<a id="canonical-fe1fc4ec53c4b91a349d259a8b4e2c69ee1fd810361ea280c1d20b44dfad863e"></a>

## Direct properties — client_side_defense / ba208314a041 / 3

- [policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6): complete subsection reference.

<a id="canonical-b7e6a2f68a0b31697adc8175d061902a7c58877d8998d83325b16a1a2781e8c8"></a>

## Next pages — client_side_defense / ba208314a041 / 4

- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3f2b96aa2f190dd702613697bc5dc00712771a3f2a168faa7efb3873d7276bb"></a>

## client_side_defense.policy — client_side_defense.policy / 123ddc31eb02 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- client_side_defense.policy

<a id="canonical-8e41f98f1fb19aeb1e6a80cd2f8c7f5851e66b642e968c56fe91cbba216490b4"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

<a id="canonical-9ca72b437418c7514b75aa2852c4a1de427755d81cd110e02b6f8c52178e5a89"></a>

## Direct properties — client_side_defense.policy / 123ddc31eb02 / 3

- [disable_js_insert](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-5836dd07a7f558503ac4af8e5a5976f2142720bbcd178e18ae5bafc6ef5101b3): complete subsection reference.

- [js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2aa33ed619429a2409ead27a7980dbf77ce7b7d686d3145d3a4b5ca3603ad586): complete subsection reference.

- [js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e): complete subsection reference.

- [js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d): complete subsection reference.

<a id="canonical-fae861119268b7ef058b84a67fa70d4a87544fd48dbd71f2c46113fe825e929e"></a>

## Next pages — client_side_defense.policy / 123ddc31eb02 / 4

- [client_side_defense.policy.disable_js_insert](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-5836dd07a7f558503ac4af8e5a5976f2142720bbcd178e18ae5bafc6ef5101b3)
- [client_side_defense.policy.js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2aa33ed619429a2409ead27a7980dbf77ce7b7d686d3145d3a4b5ca3603ad586)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5836dd07a7f558503ac4af8e5a5976f2142720bbcd178e18ae5bafc6ef5101b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2bc8d5b55d94056a3e2f746582782d2f3bfa7d76571f0080da366942a2c0a032"></a>

## client_side_defense.policy.disable_js_insert — client_side_defense.policy.disable_js_insert / 8475dd023252 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- client_side_defense.policy.disable_js_insert

<a id="canonical-6bc89be1e8c72ae66cbbd67d2cb55aef4f06888a317e9a33d89bbd05ccacd82a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-9bd21b9362f26a029ea445d91d8c88bc248b5eea1e75e88cd2043423b156ea11"></a>

## Direct properties — client_side_defense.policy.disable_js_insert / 8475dd023252 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-947849b91ab9f08820a9b0ccc7a3736b827ee61d26bc64d8f65615f1477e5ae3"></a>

## Next pages — client_side_defense.policy.disable_js_insert / 8475dd023252 / 4

- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2aa33ed619429a2409ead27a7980dbf77ce7b7d686d3145d3a4b5ca3603ad586"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-037771fb31a337eb1a255d0af3566bbfaac1583bb09489099d7cdaeb1fa2503e"></a>

## client_side_defense.policy.js_insert_all_pages — client_side_defense.policy.js_insert_all_pages / 90ddeb957698 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-34b4f5b33901c2ea3f78fe87425910e356e28ca5143ba91246f5d11bdcb8a92c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for js insert all pages.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-cfe3b97acb2f3b7d8c0f596531ec151401d59c6dae8ea517837d67abb4d35f1f"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages / 90ddeb957698 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-50ea96e70d6d193b9648bb270220a4a86481b4ceef2cd6c333e508c7d9854417"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages / 90ddeb957698 / 4

- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da91335ec9d4059448677f3adc832fe779592580bfd22e454c9075e4e43e7abd"></a>

## client_side_defense.policy.js_insert_all_pages_except — client_side_defense.policy.js_insert_all_pages_except / d2f5acb8fe34 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-5270e8d6509388fa0abab06f0c0272d4f49817909b076126ab816ae7f4d1d489"></a>

Type: `"single"`. Computed.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

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

<a id="canonical-651db10a0a087c42a32303c9c4cab1fe1aaad8a864074f6c7607a1a9540ec25a"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except / d2f5acb8fe34 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1): complete subsection reference.

<a id="canonical-6a3a01950b5ced329aa7b7c1cf7f5387a3b83e50a933538cf5c804a0c308114b"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except / d2f5acb8fe34 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e39e6963fccba7e4499c95e73b6553baa909a22ad18d05f90c0f96fc2357277d"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 4167b77f257e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-d805133c36f5d94c10268e4a645ed16e4b9b2827fe5a5f9e2329c0bf2d7d6976"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-60bd444122fcbbfdfd1b99c74334e5e4d796a5c962d14c2398715eb8b1a8f9ba"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 4167b77f257e / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-94855ad448a88cf6d1f0966ef2b430ce4d6c0a1201f1355d72a49e36761e5025): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-6c9268caa401719da63201cc3f36398f8b0737df3035d8f4afbc5fedfac88dfe): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a1a1077e5dd058a579f33fbe1baa54d60f5f1536dd90f8fcea6ad3dd90d0d2c9): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-22b217ef19ca444578a1a62236f0bd2cd670a7551ae4e5f8f56953655a1dfe1f): complete subsection reference.

<a id="canonical-d27ddf4842542b8a966df2bc8f905a53834430042ee23458bd40c02978b0c2b4"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 4167b77f257e / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-94855ad448a88cf6d1f0966ef2b430ce4d6c0a1201f1355d72a49e36761e5025)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-6c9268caa401719da63201cc3f36398f8b0737df3035d8f4afbc5fedfac88dfe)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a1a1077e5dd058a579f33fbe1baa54d60f5f1536dd90f8fcea6ad3dd90d0d2c9)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-22b217ef19ca444578a1a62236f0bd2cd670a7551ae4e5f8f56953655a1dfe1f)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-94855ad448a88cf6d1f0966ef2b430ce4d6c0a1201f1355d72a49e36761e5025"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-579b473e17a76028c58cd12a979bd0352d79d0182bbf0c61062f0ed0ae2d55f9"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / b4767c51b8fc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-b66c9b99c408fcea943644a18b1f0a5139a6ba347f722e727fb8786d8ecc3597"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-57e959ea92d66248e18c13bd1c51d706c8efc61592c79b383f01548d918e1ea1"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / b4767c51b8fc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7353ff12ea2724830156832d3356639c9897c6f746747954cfc4f5aecf28a0aa"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / b4767c51b8fc / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6c9268caa401719da63201cc3f36398f8b0737df3035d8f4afbc5fedfac88dfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c02d9a1f1d76ac5e4fa8d331a996e8e864cd56cdb55100527e4bd6c55a38ab7c"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee77103a2c1b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-07eb83da01a8e6444066fa5283ea1bd077da4ac0e9b1afecba9e456515c741d0"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-0611d9423902358ed7e73695af925dc849574cd968660039a1f922af61fcc3a5"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee77103a2c1b / 3

<a id="canonical-94217daf239ab80e2dfcf992f8c000c9719860df6c46823092c0376a8efa2e29"></a>

<a id="canonical-e3113b388ce9ea34c49f9d3aaaf0d6630772d4bfcbd0450671d1340bca05f9bb"></a>

## exact_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee77103a2c1b / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0d3ec07a669c8658d83bdbd946e89d28be54512ee7921fbdc880a8283679eb89"></a>

<a id="canonical-030537982e279857596bc7f762b2065ae29b32f678fc1dcc08b76f093da3e691"></a>

## regex_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee77103a2c1b / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-66ee4b5662e8c99496205875ec2cb60cbc4847a1fcb17c5005d4db01d3235bb4"></a>

<a id="canonical-fbade6d25f76b94bd280e454edc856abaf2f8e0623c699c5a4ba90a170cca64b"></a>

## suffix_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee77103a2c1b / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-9b8430adb3403591a4dc17b22671aa9c95f3e3c3fbd8a45597ebf30641981fcd"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / ee77103a2c1b / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a1a1077e5dd058a579f33fbe1baa54d60f5f1536dd90f8fcea6ad3dd90d0d2c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fcf8501a44c1ba29869550d91f23349511c324e7dbd933cc45130ffbb2ce22d"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 3db4f8a64750 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-029f45fc654e63874cfcb3a3619e6fa6153ca52d9983b63f4b0fba29d03e4df2"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-d9a259321d7f9bd2b9cd5e0c4a9a607b8b112a41a2c6a933b1c42c6ac1eb8033"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 3db4f8a64750 / 3

<a id="canonical-6423a58e79786a46c6a55598784adf4640c6069b5f64cf9db3176340497ac37a"></a>

<a id="canonical-3d932d06fbb0d7dfcef7c19d49fed1f79c67664ff55c2fb807b43f1b8a43f5d7"></a>

## description_spec property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 3db4f8a64750 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-6223b40e846383caf4c6104104083364c4e93fa578859aa05652f5d50c574c77"></a>

<a id="canonical-2a846d3ccc605c55233d4f3a0abba5b26193bbe6c5b962980ff7704910958724"></a>

## name property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 3db4f8a64750 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-752119b08be30c40961833f128f607b4c26ebd588beac53b769705763f5ff576"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 3db4f8a64750 / 6

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-22b217ef19ca444578a1a62236f0bd2cd670a7551ae4e5f8f56953655a1dfe1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-733d7ab6d894cac0db52b549a8e72b0f0d76c9f00a3a8a13e715aa314cd116e5"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.path — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / ccdf36009701 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eba0239b3793eca6d9b31dd704bf113c035ac3b7bf49cd6dd95a9a2cdcf912e)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-87078b73fc25b8c2e1475ccc8919cb1dd77b27014acc3b6b3b576fd9099d3382"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-0ba30e433a5d0a4f77cc7fb36cf81998ca7aeb89fa94f464642ec0ef281e95b5"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / ccdf36009701 / 3

<a id="canonical-83754da41a3195eb9d5d68a25a9768c36b195793f2ddabddb03d8cb54b726c85"></a>

<a id="canonical-db3e6f633464b5b5c0ba5b57fcb54141da3f94c8072abbdb02e4cb76d554979d"></a>

## path property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / ccdf36009701 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-683d52738370544964fed0da3a301bf640d189041c3dbecf08aa6d0c21ec7464"></a>

<a id="canonical-41ecce49e4e63ff31805896de03ac2a2343796bb691d84e5b7e568e11541f0e8"></a>

## prefix property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / ccdf36009701 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-647c8ab51b25d9d82f3de717ab7ddac7da0b5cad6ebb431180589c1f8412f2fb"></a>

<a id="canonical-76097f9b3c9fa1d18a81329c5faa414b8c04621fb9d2339bb674b4be9d178247"></a>

## regex property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / ccdf36009701 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0711e57d3c77331250d9d44e49308d8ac36c262b5f0eec80a7120118afddffdd"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / ccdf36009701 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1f1cd264a5a96eed55349dbef3f3bec569a1e91267b52c2f7a03e1d0cdbe76f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2db3a0b76e7d6e898999d056bf7b5771e4708a3b96c070bff0c1b4d2a0a80e5a"></a>

## client_side_defense.policy.js_insertion_rules — client_side_defense.policy.js_insertion_rules / 0fdfbe178610 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-4f406e841a5ef03480e4546087fbcee66b6f102099accded6d3144d571d99999"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

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

<a id="canonical-0cf54e5903b7f2908aa96c9dcd82905d4a2b4953f542287556d421b3b3549b9a"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules / 0fdfbe178610 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a): complete subsection reference.

- [rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1): complete subsection reference.

<a id="canonical-1707d1d19ee134e312bfc72c8bd574eefa2804a99f3e0c70007ad3c1a22b9e8c"></a>

## Next pages — client_side_defense.policy.js_insertion_rules / 0fdfbe178610 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f537828733e3f0741ec5f5cf54028f0e27627bfbea4d557d397dc82be8f6ca64"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list — client_side_defense.policy.js_insertion_rules.exclude_list / 211fbc0c6399 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-a572b20c6d194ceadaff6c74e2a4c5fafea91f06457c1757994255717eab9d5f"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d12b40f16cbe390c961b7c04c30d0c69b1320ce6b97eef24c16f17b010808215"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list / 211fbc0c6399 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-c6ace50f5c616c196f66adb72fa8a462a0408f8d727e8a6052995964b9cf9e9b): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-dd678662e99eba003d29aaecab51371751b5ca5b7b461b910b966018a7e65a3b): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-7eafa2d1ef4f761746131b2657d86ceba4e82d7571aaaa16c2a2d2856d2edaf9): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-be2935546a6e2a2dcb30e431d16cd96d2f45e837991fbab11572179875b773a1): complete subsection reference.

<a id="canonical-114a79240db9aef81fb3f0a3b35018321556f437157ca9c8b7ccc0b92cec2591"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list / 211fbc0c6399 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-c6ace50f5c616c196f66adb72fa8a462a0408f8d727e8a6052995964b9cf9e9b)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-dd678662e99eba003d29aaecab51371751b5ca5b7b461b910b966018a7e65a3b)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-7eafa2d1ef4f761746131b2657d86ceba4e82d7571aaaa16c2a2d2856d2edaf9)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-be2935546a6e2a2dcb30e431d16cd96d2f45e837991fbab11572179875b773a1)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c6ace50f5c616c196f66adb72fa8a462a0408f8d727e8a6052995964b9cf9e9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b14e02c1e35e5338df0bb2fbf0f4af3fd38a211c4739350b4d9f919c39d6419b"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.any_domain — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 9d5a41ddd9c5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-722caead03a83b52fa6e12a55b822957eda615e21d7360723196399b9bf4cb4b"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-4fe3908b3fc43957382109c3b179ed6f85dfec75167f8fc8435b0f3f22bc0632"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 9d5a41ddd9c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0d371be5b66e5594095105157f30a055c2a0ee66d80befeb5c7da114c88685c"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 9d5a41ddd9c5 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dd678662e99eba003d29aaecab51371751b5ca5b7b461b910b966018a7e65a3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22f90f2a5f8a356161033cfc4c4d79d1dcee59741215c056986fccc15ea6a576"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.domain — client_side_defense.policy.js_insertion_rules.exclude_list.domain / bfcccee68857 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-8bd5fb42b2c7c0d1e3f5d6c617d22f126345583f88b932cb598b389358737cbc"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-b02d642072a66ed5cf1f03b53dc26b8c9ea1ee2c4cf55b1f248af73646ccdb22"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.domain / bfcccee68857 / 3

<a id="canonical-ea469a4030a5c3e4b5333e0e873a31b05aa2fa41e854ca8a3623c69eef050270"></a>

<a id="canonical-071de6d183148ff5fdff6f62f457659ccdea9315d608b6d00bca03a16615e902"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / bfcccee68857 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2d87e4a341b94ea242cc315de56f1fc5348ebdcf93b977a141ec8bc556ea37f8"></a>

<a id="canonical-77fa5c43384fed2af30498dada3c91c3f6f17b6a221a47bc435178db18a24cfc"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / bfcccee68857 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-f95c5e72f1aec45dce4a7d5f6b788878955fca352402e6a5aee20e60b36e9bdd"></a>

<a id="canonical-83b6d5c31c93b2ff32c04698aba9459ad2d157b6ed88ff41c20e204259fd1068"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / bfcccee68857 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-bd4ddff261f46b5cad5867a6e7144bf7960db6f73656a986ea28e5735fa738b2"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.domain / bfcccee68857 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7eafa2d1ef4f761746131b2657d86ceba4e82d7571aaaa16c2a2d2856d2edaf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-391a409b11defd8aa842ab8ca1f666f598d59192cc1615361f0d39b60b304f9c"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.metadata — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 224a8dfc70cd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-81ab8e88a3fa907be6d5a001e0c7f7c6b6f43c50d675cb709b8119068d90a03e"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-95b12addbc6a58d787b801abd9f19d2fea0a7118acc108d4db491405b901473a"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 224a8dfc70cd / 3

<a id="canonical-b431df3844239270893068cbd12108ca09587af66753988ca88dac5dbf60e841"></a>

<a id="canonical-a089a40e821f7e125f4f2f3d1488ff548f0565f5c805d855a7235ffab5620272"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 224a8dfc70cd / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-392ebbdedbafebdc7344eb7385916f89229ffb3a0e498d9c5f674208e5492699"></a>

<a id="canonical-927635bab0e4912c38ff67f7dea5b5baa72f2c56a51be3b7b30300e4d4e34791"></a>

## name property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 224a8dfc70cd / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-aff8abd7864d707fb046e7a3f90e09956936b8e5090b844fa585be68b61083d4"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 224a8dfc70cd / 6

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-be2935546a6e2a2dcb30e431d16cd96d2f45e837991fbab11572179875b773a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cec166a4d7701c413f3518e271a918dc7eaa65decfb383ba65b45bc8c67308c1"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.path — client_side_defense.policy.js_insertion_rules.exclude_list.path / cea223befdb2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-7a7943665150d91bc797e4d682fc19ddfda8c3b48fbbc396a64824076080b58b"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-c479fbd0332a07dfc995f8b3ef2cabb44cb441b5d048c59bb1fca817a18aca53"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.path / cea223befdb2 / 3

<a id="canonical-729900102e028d31f973d9d51dc599dc068c8d7ee780086a9af18390da19ae99"></a>

<a id="canonical-2fdb2dabea2264eb01a22d9f85fc36b9e3746e029361e6806382990eba165ca0"></a>

## path property — client_side_defense.policy.js_insertion_rules.exclude_list.path / cea223befdb2 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-97929c526761ec7f024d95bbe95aab8e8d2c0603190e3e57ed3fa71dc5d9fe8f"></a>

<a id="canonical-7e50b90b6c5dfeedc1f43ab0590c7ad446dfb5659d418949fbdc5383c4a1f565"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.exclude_list.path / cea223befdb2 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-83619f1a6e67607221f0c43ef784408065b4ba9d1b4c15fe64382b4773fda9e5"></a>

<a id="canonical-ca07bc7a4a263ecb9317a20f1ee01587103c5c6afaf32449b0c94c00f054b659"></a>

## regex property — client_side_defense.policy.js_insertion_rules.exclude_list.path / cea223befdb2 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-491c090dfbb5ceb78a033dfeb75eeff05f090a129bd1d29565d8828084096432"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.path / cea223befdb2 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ae4c8ea09004c9ccc2eda04d31ed4a8be51c6e9bf4def07a06ec1bc6cbe8791a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-352f605b5ad2ca3db3427ed1787ccd38f5c45c930dc27b24de02641494a1a1f1"></a>

## client_side_defense.policy.js_insertion_rules.rules — client_side_defense.policy.js_insertion_rules.rules / b979d46c5cbd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-e0b1a80ef0c7cb570d569c64338fde83cae1b0339a14e31aaa47c6ee7125293b"></a>

Type: `"list"`. Computed.

Required list of pages to insert Client-Side Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-573c5e3637d1c27677f386d38c74b72c4db3aae76199b97c0149d0ce333cd2ea"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules / b979d46c5cbd / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4ebf4c2319349a502070451708762fc115ebee2a5385050a04f91643ca634eea): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3db8e9c8f88da70121d5fc0db983fd4a2a5325028720d923a7f0669bb532482a): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3a5944f69e341ef51083cf5083d02ea98c32264e2ea7bb14fbdc88d2571fa7c5): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e5a547f00445eab773a0a8543b1d784c0f65d62102b56075d519020502670460): complete subsection reference.

<a id="canonical-6047725abb7492dc9f2ceb828fd380e767d6e61e89e50cd27c7aa54f37689ddf"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules / b979d46c5cbd / 4

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4ebf4c2319349a502070451708762fc115ebee2a5385050a04f91643ca634eea)
- [client_side_defense.policy.js_insertion_rules.rules.domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3db8e9c8f88da70121d5fc0db983fd4a2a5325028720d923a7f0669bb532482a)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3a5944f69e341ef51083cf5083d02ea98c32264e2ea7bb14fbdc88d2571fa7c5)
- [client_side_defense.policy.js_insertion_rules.rules.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e5a547f00445eab773a0a8543b1d784c0f65d62102b56075d519020502670460)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4ebf4c2319349a502070451708762fc115ebee2a5385050a04f91643ca634eea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e510f8e44a864b96441b35b5405d85e27c7d1de6d6bbc5f3d971bdff1fbbbcc1"></a>

## client_side_defense.policy.js_insertion_rules.rules.any_domain — client_side_defense.policy.js_insertion_rules.rules.any_domain / 48ed2ce95d9c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-4949e7a947fcda86c3404dc1b112a4818ac967270fd88c53ea16b8f4cb168398"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2f399e70a1df9157d11f020af5d639518d9085602ba4e8f2e16a4f28960ff1ca"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.any_domain / 48ed2ce95d9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4cf3f31c0a35b5e5f527ae395e92637b79c62a10c235be9518e495b29eed9655"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.any_domain / 48ed2ce95d9c / 4

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3db8e9c8f88da70121d5fc0db983fd4a2a5325028720d923a7f0669bb532482a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47b6412ff26e80af300513338083db199ed83dac7cea85a8329783b9a511368e"></a>

## client_side_defense.policy.js_insertion_rules.rules.domain — client_side_defense.policy.js_insertion_rules.rules.domain / b7e7accf5172 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-fb7fcc6221b9276e2ffb3c2cb5ee73fbf7be19f332226f614bd6ab0488f5fc1f"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-c63f46b2006a56d79aee0f07a6b9c9254f0c15592ff10bb2219d74d727ea5ed3"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.domain / b7e7accf5172 / 3

<a id="canonical-7c40a2d4b2e344e54a4b6822e4fcf076c182470c3b64bd3fcf034617c975c289"></a>

<a id="canonical-031f32ed2ae766ded468efaa09e27097ea8aae56b376ef9f766575482616fd47"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.rules.domain / b7e7accf5172 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-28ad67f9fde61d1845f04bd5a351d37043e5035e5bc7cf406f32978336d0671f"></a>

<a id="canonical-31df3ee3a2a46db17d3d1ca5ec16573b4ec9a09fc95b81c88bba3ffacf9dac1e"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.rules.domain / b7e7accf5172 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-9b130363aab38a422e3d27f9c239e962b6559f26fffda5837d3f5d5caab7f23a"></a>

<a id="canonical-675b3285d0c09a3e553b2cb76806a4d1347365fced43c8d6f6909b31c1d334de"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.rules.domain / b7e7accf5172 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4ffe1880ee5906ba27f262d3ad9fb2bb99d267b315d6f7f7cf92b6ba5a3f3de7"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.domain / b7e7accf5172 / 7

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3a5944f69e341ef51083cf5083d02ea98c32264e2ea7bb14fbdc88d2571fa7c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32d6c75439a38d5a883a229be06009aebf10ef99285bedc5cdd3151a5949ce7e"></a>

## client_side_defense.policy.js_insertion_rules.rules.metadata — client_side_defense.policy.js_insertion_rules.rules.metadata / f88988949ce1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-0425d0f577225b9f125b5cb87c1ccf2d6b93b3b3c6e99e2fbc6f6450ff3fd5ea"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-af709cacb1964bc4b88c54231330050c02be329da3d9e9f0d4429e7141413939"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.metadata / f88988949ce1 / 3

<a id="canonical-1c0dc3b548169da11bd96370311c18c39ab8d603a915010b8937c566241531fa"></a>

<a id="canonical-880db3ba99f32d98fd1112a8619916efda13d4496e10e9992f010b92479ccd7e"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.rules.metadata / f88988949ce1 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-049bce61ba2eec05837d9069434fcd310a4a4721d5c2c5e432d162c9dfd58195"></a>

<a id="canonical-5f46fb2f76d984ad191e1e748d5cb0ba9376c5681401e39a3aeca0872ddb8808"></a>

## name property — client_side_defense.policy.js_insertion_rules.rules.metadata / f88988949ce1 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-7776ce6c131029100b5567887797e581cbd4633423017ecc11dc9c1a3ecfd1c0"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.metadata / f88988949ce1 / 6

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e5a547f00445eab773a0a8543b1d784c0f65d62102b56075d519020502670460"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8ea7ecf877e078c5e085abf9ede08d605eb7a15da12d3d77908b648850d9142"></a>

## client_side_defense.policy.js_insertion_rules.rules.path — client_side_defense.policy.js_insertion_rules.rules.path / 968f3386e1de / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4d59a2b9c6e40fcbae06219e397d5552634502ae27be771799d796ae596e512d)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-d8815d8ba61df5d381314b1b3eefa177d7d12308c58a49d2b945795d05e4e9f6)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e731ca51d2a91fc2ff8b033434442cf373b9570c672ae833b3d0b5ca7eab246d)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-d4c81d0224b6f8d047e262425691c498a46cb4ada539ee5ab5b3b4185024bc0b"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-75f865d8401733bb0967561cd4ea082ff68e8d5cccae10755b0012ad28f74a2e"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.path / 968f3386e1de / 3

<a id="canonical-dff1fd5f00043f1a587f1e3df7304314b35b744cff4d47666b5ab489919cfb8b"></a>

<a id="canonical-3e8885fe02f3f46c9daca5f40d6a5c72fdee2bfb2ffe2aee5b9486bbb7fea9e8"></a>

## path property — client_side_defense.policy.js_insertion_rules.rules.path / 968f3386e1de / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-c5bd7b41938b77126afe2e2a5cf113dc6a2f1002e7983c9140a51f46c17ec0b0"></a>

<a id="canonical-c52d4de93e1fc9be244c821055861ffbf28dec0edad731a5e84929174f728cb7"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.rules.path / 968f3386e1de / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-a0ec4222e0fa666a7301249b0f8582140553f449707361db8cc95d43e3d25f00"></a>

<a id="canonical-89cd30d20fbf9b2693a7a64aef367ebd6ccc4833adaa5f6df367414bf76f3f71"></a>

## regex property — client_side_defense.policy.js_insertion_rules.rules.path / 968f3386e1de / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-f387d5fdb57e8914eb9d52c26588f49e1c1ecda53a91f12ed02a480c2bee04fd"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.path / 968f3386e1de / 7

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ac85ce92b74c10a6d3b5d89bf67a7fa73203ed3fd43adf4b64e7992ff91a08d1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-fc417be8204c4b8728a483a1863650a03902296f2284d0eebabd5417667dd223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcb46b103407d7d0e1f9d6a9fda42e5b1b9b83935749dd120c8ed54472d4c179"></a>

## cors_policy — cors_policy / 9ba87595d088 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- cors_policy

<a id="canonical-495ea2c7969c36dcb4531446ac6f0162ddb25fd63e8bce1a6f6c0ba8f2f0a733"></a>

Type: `"single"`. Computed.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.html Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

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

<a id="canonical-4978ba8f5d68d88369a13e8494ba5d4bbea1696acab8b348eb6ed3f1028f1261"></a>

## Direct properties — cors_policy / 9ba87595d088 / 3

<a id="canonical-2bf23b6289e4e60aa9e361ba4d3b69a7f54f6f5f80b1ef9d9176244e5c83c719"></a>

<a id="canonical-6d1ae1041368f55fe0d3d0746d6965486bbb80ef9ae1797aec263fbe47691bec"></a>

## allow_credentials property — cors_policy / 9ba87595d088 / 4

Type: `"bool"`. Computed.

Specifies whether the resource allows credentials.

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

<a id="canonical-cd0e5e960e9669cd850da19b34bceaa3126e292649e07df02866c34b5d4379bc"></a>

<a id="canonical-7e31be0f0f1500477482aace43c96996555e3775942bfde8cfb9687e5e2e19cb"></a>

## allow_headers property — cors_policy / 9ba87595d088 / 5

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-headers header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9e5d974679da261de32778443704c41d3e22b3b8264f26342b9ecfbcbff4565b"></a>

<a id="canonical-c8dbb0a0f69fafa091aa33521cf174042b26d4cc267dc3ae6f7bb413c3a29eab"></a>

## allow_methods property — cors_policy / 9ba87595d088 / 6

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-methods header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-9fcf1996ae330589b419aa23c3d706fb9d699bd7b20b447cd060143b42e3ac4c"></a>

<a id="canonical-77260868c637de8b2aa8f39d6d4e804610dee9d17978f16db00eeb9e237fe1d9"></a>

## allow_origin property — cors_policy / 9ba87595d088 / 7

Type: `["list", "string"]`. Computed.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8bb2329c2d0124daa2d656d2db9bac4b99b3e3b2ce68e16443a1a824dd5f28e5"></a>

<a id="canonical-f3aa4bea08de8bb6ce90fd52603101eb93df587a1ef791e20f47261b183af5e8"></a>

## allow_origin_regex property — cors_policy / 9ba87595d088 / 8

Type: `["list", "string"]`. Computed.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8bd21bfa8f68e88b0f15da4aa5214efee74abd28cd873d7acedf44df95f753fe"></a>

<a id="canonical-36b4604755af1cb090da0e323f5b6ecf49d4e8e8bd0ae8b546cd9b0e75ac9ac6"></a>

## disabled property — cors_policy / 9ba87595d088 / 9

Type: `"bool"`. Computed.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-99220b1a50d08a1b814e74df5b8c9f734449bbef1d67518e1e8a159318bb048c"></a>

<a id="canonical-c9ea8eeb634f3f537e33c570e161415906ea900e7e1ca266190b5f47755eedc4"></a>

## expose_headers property — cors_policy / 9ba87595d088 / 10

Type: `"string"`. Computed.

Specifies the content for the access-control-expose-headers header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3f23c154a37b0efdfe5d934819fd1e68400229627b0157ff65c4d19889a4fb6a"></a>

<a id="canonical-976860b6b473027b5389aeaa0edc1d721ae613cbba24acc5e5fea84dabde5228"></a>

## maximum_age property — cors_policy / 9ba87595d088 / 11

Type: `"number"`. Computed.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-0ee6cf4ae63b9e15dfef9bb0c38b69650343ddb5a8b755669d9a70bee0e7a8bb"></a>

## Next pages — cors_policy / 9ba87595d088 / 12

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c2ad8bfb261be3f69e4be05074342d8460e876910ea3624860a05fbcf0b1e31"></a>

## csrf_policy — csrf_policy / e3728054de35 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- csrf_policy

<a id="canonical-8e0b5d87005a8d121ff8719feb17a021fe003fdaa122ff0249b118358bb423dd"></a>

Type: `"single"`. Computed.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

<a id="canonical-1018ec66f14a0f80d5acf8956993abc414d51e44b4f29691353ff2b2130ebba4"></a>

## Direct properties — csrf_policy / e3728054de35 / 3

- [all_load_balancer_domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2cf8f3f1b24adf364705575d680007e134b824b1c96f909dc679f9651e3960f6): complete subsection reference.

- [custom_domain_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dd2a9e17a1ab0978a31be668a514ceacbcb2e35d1490fc9a2851c729a5155aa): complete subsection reference.

- [disabled](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e54caed09a325e7e71f9429fa9ad6f1746ba320bd9e2b2f80821ae39ee3f5846): complete subsection reference.

<a id="canonical-9730ef3ce7ea9f76dfb5634cb99b049413bdd4f73d417be645fcd0d6f74ecd87"></a>

## Next pages — csrf_policy / e3728054de35 / 4

- [csrf_policy.all_load_balancer_domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2cf8f3f1b24adf364705575d680007e134b824b1c96f909dc679f9651e3960f6)
- [csrf_policy.custom_domain_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dd2a9e17a1ab0978a31be668a514ceacbcb2e35d1490fc9a2851c729a5155aa)
- [csrf_policy.disabled](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e54caed09a325e7e71f9429fa9ad6f1746ba320bd9e2b2f80821ae39ee3f5846)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2cf8f3f1b24adf364705575d680007e134b824b1c96f909dc679f9651e3960f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cb069a1bb47e043b7de07c35f02a714b5e77dfd21f1d745f114e153b32937c3"></a>

## csrf_policy.all_load_balancer_domains — csrf_policy.all_load_balancer_domains / a75a337f7ad2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c)
- csrf_policy.all_load_balancer_domains

<a id="canonical-51b6779b736d121d507947d1ac132b3b7983f6706874bc8457766e4f86bc5eab"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all load balancer domains.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-79a002db290e2671e04b41c32e702952d631e08a948e5a625f8a1505ef588fe4"></a>

## Direct properties — csrf_policy.all_load_balancer_domains / a75a337f7ad2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-999c47dc8793063d30a0ccb6a0365b3849859e462b3327c7fec430c166619278"></a>

## Next pages — csrf_policy.all_load_balancer_domains / a75a337f7ad2 / 4

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2dd2a9e17a1ab0978a31be668a514ceacbcb2e35d1490fc9a2851c729a5155aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1c8e4e6c1c2961160ed752d4d16dfe7358c6f74f89d5f4192c958e7a7f61174"></a>

## csrf_policy.custom_domain_list — csrf_policy.custom_domain_list / e0932e4d0755 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c)
- csrf_policy.custom_domain_list

<a id="canonical-f8f02e03caaeb96ddcbce2b7a527087fa86420b24b32c5dfb328b5bee8d4baed"></a>

Type: `"single"`. Computed.

List of domain names used for Host header matching.

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

<a id="canonical-4e1b68db63d279e58b9b61f186536a5fdd65f2285bbfbc884e7b5a0e37c901e7"></a>

## Direct properties — csrf_policy.custom_domain_list / e0932e4d0755 / 3

<a id="canonical-7793a79b67d3fe39b632b57f6fae287d66773177afbf723a33a5cec44dec9948"></a>

<a id="canonical-635fac55b971a7793639b8b6c750ff61d15640d6eb701912b5a496e83536dcb9"></a>

## domains property — csrf_policy.custom_domain_list / e0932e4d0755 / 4

Type: `["list", "string"]`. Computed.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

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

<a id="canonical-d5f26d1e8ae07108df9c04d6a6153e90d402fe691033004bdf7728427ff04edb"></a>

## Next pages — csrf_policy.custom_domain_list / e0932e4d0755 / 5

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e54caed09a325e7e71f9429fa9ad6f1746ba320bd9e2b2f80821ae39ee3f5846"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e78326412ba11fc2c922d9371e540974f302a32b3822342f489e707f8d875352"></a>

## csrf_policy.disabled — csrf_policy.disabled / 84ba145f585d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c)
- csrf_policy.disabled

<a id="canonical-ac45e0afcab4ae0e77344bdd1b229bd7d4fb34bbe1bc09ede277698fd63a5b0d"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0a5a90bb6ccd269d1a56c52c3871f96c0b2652d95601a0238e298ad9a16b5e52"></a>

## Direct properties — csrf_policy.disabled / 84ba145f585d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ff3543b0b81d23889ee84437be8e5f2e3da7894922bc6ec3dd812e7dcf0d7c1"></a>

## Next pages — csrf_policy.disabled / 84ba145f585d / 4

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-75eebc6e30f92011b907ed5c9d3ab2a72625997eccf90b773ec6b505f338396c)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-06d10c59251d252d7fce9adb3b74d3979ae27b09a55d4941edd7d01e37b4af78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-569f96d36577da70d8c94fb8f1599e1a9ce8ab61e117a97c5dc79fe8519dff14"></a>

## custom_cache_rule — custom_cache_rule / 07ee514ebccf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- custom_cache_rule

<a id="canonical-82be8cb205b9e2c4f7bd39d7eed62420504f282d1911bf733c267e7eec277fed"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

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

<a id="canonical-3538be3e2979f22218787f420a1b4c97c8dcbd2f984291c2e38c666946d30556"></a>

## Direct properties — custom_cache_rule / 07ee514ebccf / 3

- [cdn_cache_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-89fb7415fedf7e76b975205a41cbacd69121de0afdbd23aed7bb54634627d0e3): complete subsection reference.

<a id="canonical-e8ca6a705f889b59e293efba41e5e252ec04a3e645a6a3bccfa521aa78cd2acd"></a>

## Next pages — custom_cache_rule / 07ee514ebccf / 4

- [custom_cache_rule.cdn_cache_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-89fb7415fedf7e76b975205a41cbacd69121de0afdbd23aed7bb54634627d0e3)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-89fb7415fedf7e76b975205a41cbacd69121de0afdbd23aed7bb54634627d0e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee6748d30e45ecadb5e64516f3943d8d032030721356d0d459398961ca6a04ac"></a>

## custom_cache_rule.cdn_cache_rules — custom_cache_rule.cdn_cache_rules / ed4419d32723 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [custom_cache_rule](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-06d10c59251d252d7fce9adb3b74d3979ae27b09a55d4941edd7d01e37b4af78)
- custom_cache_rule.cdn_cache_rules

<a id="canonical-adf008583fbc03fb91cc0736d799544076a6e217b87b8feb008d804962f455ee"></a>

Type: `"list"`. Computed.

Reference to CDN Cache Rule configuration object.

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

<a id="canonical-aa40550030742570d1fee9b2e854575c3400a5ca758143e6bfa966e296087383"></a>

## Direct properties — custom_cache_rule.cdn_cache_rules / ed4419d32723 / 3

<a id="canonical-c6356092b9824a51ee826e7c211c1bf13589cdd754d0ee90a49a37bcd6d55a58"></a>

<a id="canonical-f9110bf8c5f14014dd177c8c8d12168df736ddf4ada7116a15491d5f6297e607"></a>

## name property — custom_cache_rule.cdn_cache_rules / ed4419d32723 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-5d822b2210cf30fdfa669e221e4fce8c746061c6a9799321618a8351d081dd52"></a>

<a id="canonical-21a58476cc5aefe956468b1ff94f5dbdd0f0f4e740718d905a69c2451258e1b7"></a>

## namespace property — custom_cache_rule.cdn_cache_rules / ed4419d32723 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-dd56eccf7ab2e83cf9997d62efd108e1d4b000e2b0b62d4bd4ed1c9a34077f7e"></a>

<a id="canonical-0166135911f42ba7864d09a8282c729b6f222fd0405875e92bef03b69288ef01"></a>

## tenant property — custom_cache_rule.cdn_cache_rules / ed4419d32723 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-40059c5b56fd3ac0224e4a40ac429d418ac84a40117fb93adef4693f247f5126"></a>

## Next pages — custom_cache_rule.cdn_cache_rules / ed4419d32723 / 7

- [custom_cache_rule](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-06d10c59251d252d7fce9adb3b74d3979ae27b09a55d4941edd7d01e37b4af78)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-008a96df9e27b689fd5651607c497bf010d171adad533ccac8548372591a6077"></a>

## data_guard_rules — data_guard_rules / 57a989bc7433 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- data_guard_rules

<a id="canonical-a4c502e59f65774c9dfdc20eeb6138b2d86f9a4865538f3bff1d06529dfd73ff"></a>

Type: `"list"`. Computed.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-f3bdae70e9e71333cdc5fb46f6cba0669757ae3a844e93f76ce54d0606ac3058"></a>

## Direct properties — data_guard_rules / 57a989bc7433 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-eee5918dade18681cc2539b58e93b41e551768f7b4c0dc407d1a42b69bf08c48): complete subsection reference.

- [apply_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2d816f84c54cf771a41ef170c0e574d9bda69c62fdef9e6a6456d52f315f0529): complete subsection reference.

<a id="canonical-9d0845a4f6b82f139b60fe1390c387e31e567896d3400484e18c95298e114266"></a>

<a id="canonical-5eb3d8d3eb9c1e3b6567a4632efe52f04ebf05a1f36e93b7e2aa69c82bed734d"></a>

## exact_value property — data_guard_rules / 57a989bc7433 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-6d34e58f1ec43d5063e7ecf2555ba00bc82a3f96b7e8cd24f0250b2200a70811): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30648fe6c33ee55cf713d79b064e8208a7387d10c9bbf038aaf433bb6f63ad72): complete subsection reference.

- [skip_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0e7d8041e88130207e149503c31d2e6d63f6110dc670a44d25a972d682a7dde3): complete subsection reference.

<a id="canonical-f449d52cd3f89729277017fac1c3fff81c04e2237bd921805b5fac7be5129b15"></a>

<a id="canonical-bfd71162f828cac9aa14110c2d93b28b04899764c3d6d26f62e689f9e97bbbbc"></a>

## suffix_value property — data_guard_rules / 57a989bc7433 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-b46c8aa52e93d3b60cd14137f85dd302e05755707f1825c55e99edac1d796495"></a>

## Next pages — data_guard_rules / 57a989bc7433 / 6

- [data_guard_rules.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-eee5918dade18681cc2539b58e93b41e551768f7b4c0dc407d1a42b69bf08c48)
- [data_guard_rules.apply_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2d816f84c54cf771a41ef170c0e574d9bda69c62fdef9e6a6456d52f315f0529)
- [data_guard_rules.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-6d34e58f1ec43d5063e7ecf2555ba00bc82a3f96b7e8cd24f0250b2200a70811)
- [data_guard_rules.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30648fe6c33ee55cf713d79b064e8208a7387d10c9bbf038aaf433bb6f63ad72)
- [data_guard_rules.skip_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0e7d8041e88130207e149503c31d2e6d63f6110dc670a44d25a972d682a7dde3)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-eee5918dade18681cc2539b58e93b41e551768f7b4c0dc407d1a42b69bf08c48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ca6e24431b614dcdbd664a72023af6a23cfe640f1c7906a26c671e5d1415069"></a>

## data_guard_rules.any_domain — data_guard_rules.any_domain / 6291ef72a2a7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- data_guard_rules.any_domain

<a id="canonical-49869c7dd1cb15f12f901a2d445ba883446a57bffed26b2ab7fae3a7408b1e5e"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-9631d0f2dd403b11015a5eb580b707964b0c73a7ef7c25721c2f1427284bec4c"></a>

## Direct properties — data_guard_rules.any_domain / 6291ef72a2a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4eead92a5c944c99cc0d7b442ae10912f3515bd9cc651c34a881ce398ddac147"></a>

## Next pages — data_guard_rules.any_domain / 6291ef72a2a7 / 4

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2d816f84c54cf771a41ef170c0e574d9bda69c62fdef9e6a6456d52f315f0529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab156728eb51ccc4fa1594707d1e29ff65bcc7ebf611dcb9f722bb33ba82ae5a"></a>

## data_guard_rules.apply_data_guard — data_guard_rules.apply_data_guard / 4f5b688d0c39 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- data_guard_rules.apply_data_guard

<a id="canonical-fa9193fa1f4b4f6b83edbe375505f3595cf2db87eb0132a38a4a0bef5504f2de"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-f087edc7abae9eb220c3a4dec4bc4f15b995271f955253ad40ad0cd4b12b3d27"></a>

## Direct properties — data_guard_rules.apply_data_guard / 4f5b688d0c39 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3ca485bbf95ee31b9054564ee204e22c6e227673617d5200d4c9b84852395f00"></a>

## Next pages — data_guard_rules.apply_data_guard / 4f5b688d0c39 / 4

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6d34e58f1ec43d5063e7ecf2555ba00bc82a3f96b7e8cd24f0250b2200a70811"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6908186b7da78c905b565a285aea80b0524c24f24f3cb7caa006448b0f42952"></a>

## data_guard_rules.metadata — data_guard_rules.metadata / 11b54b086cdb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- data_guard_rules.metadata

<a id="canonical-61290181765c52e21ba4d4ec00db56e24c7cbcc6ffd699207e9a52e8acf773b4"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-079f0cf1f89b0bdde0a8175444e9e1c1acde1d47fab80922dc05c357d2f9d8e5"></a>

## Direct properties — data_guard_rules.metadata / 11b54b086cdb / 3

<a id="canonical-307eb80b0364c78f9f44631e351edb99793928b28c16cb87209efb2e059829e3"></a>

<a id="canonical-34b2d27f7363ff4ed8085dae4cf4206089b3a266cb4903477a2360a3c703e09a"></a>

## description_spec property — data_guard_rules.metadata / 11b54b086cdb / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-d3cb07192779fe7cbed49374599da7060928b367033ea752cfd3155bacf124ac"></a>

<a id="canonical-9864f252960dc37961f7c71781236ca00e400b8ff5fbf201188b30ada6731cbb"></a>

## name property — data_guard_rules.metadata / 11b54b086cdb / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-13edd34d95bf3f02686a7c319f002c36e6438d1ba159b33df8649659a2cdb026"></a>

## Next pages — data_guard_rules.metadata / 11b54b086cdb / 6

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-30648fe6c33ee55cf713d79b064e8208a7387d10c9bbf038aaf433bb6f63ad72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7190baa21f61a5e1df35ec0b21d662af41f6fb6122ba90eff78eb8e51a3596f8"></a>

## data_guard_rules.path — data_guard_rules.path / 2fa8b69275c0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- data_guard_rules.path

<a id="canonical-80b392081148b28aea1bfb183032f0a28f6e50531b8973838741cd64b0a50d68"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-b52319b3e7fcc8337ac02f234538c5a48f87a9f67102baef4ae4fe6158462479"></a>

## Direct properties — data_guard_rules.path / 2fa8b69275c0 / 3

<a id="canonical-d0f3d8904756af06591494adf2874d7c6e0a3725030e9266aa697e9927c8b318"></a>

<a id="canonical-4d50338e1a01a5178982b0799cb6127ebbdfa94cd0bb6bdb0afd49227eeb4d8d"></a>

## path property — data_guard_rules.path / 2fa8b69275c0 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-d97d71d60e03248fd8f2b1735c8720ff39251c4ec31771ce76a698a6efded858"></a>

<a id="canonical-57d56030a9efbd3a5b97cc90aa905ea35fca75b2ec6ed2d7b893f74a52ec8db0"></a>

## prefix property — data_guard_rules.path / 2fa8b69275c0 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1f0968ad2114c059c9e47c7ecc6ef008a6bf9ec98178cb3c60704c7b380f9a11"></a>

<a id="canonical-7cc084f74196abb8e413f064b5d4cf2d503b91a4dc14449a40b83753d6343a4a"></a>

## regex property — data_guard_rules.path / 2fa8b69275c0 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-e2e07a9608bb1ed837badc62a807ce1c6a86b45af7509654cf56a1a531e4eafb"></a>

## Next pages — data_guard_rules.path / 2fa8b69275c0 / 7

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-0e7d8041e88130207e149503c31d2e6d63f6110dc670a44d25a972d682a7dde3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e4cfa38af40f9c48e81d2bc68f722662b48e17098709047be53ef4647442f23"></a>

## data_guard_rules.skip_data_guard — data_guard_rules.skip_data_guard / c09ef45e9282 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- data_guard_rules.skip_data_guard

<a id="canonical-4079750b0859295a7e1c4a024baace511d1fd07d3d282fe0187c89cd20c85a34"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-59d7db9141f278bf1618cb1a1dc84ec9c431a9299981374edb79a95c31c2dfb7"></a>

## Direct properties — data_guard_rules.skip_data_guard / c09ef45e9282 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e1702ce1a06468647fc9e28dc6f22604a7424f30d2b484f0ea4378216c68a35"></a>

## Next pages — data_guard_rules.skip_data_guard / c09ef45e9282 / 4

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-ee57a81a6530732d8c51e83ef1e9b9b1d71709c8413ab7557b32a271eb2a7e65)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3afd104f2fc4283c79cc1d5b509f0a54b6485b011bb3c3d8de02905be5857dfd"></a>

## ddos_mitigation_rules — ddos_mitigation_rules / 4a7965fb5ffe / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- ddos_mitigation_rules

<a id="canonical-13d69e243fc77ab716fc5f16eff85cf3ec62fdb9d222722318ad3d9b377d7bca"></a>

Type: `"list"`. Computed.

Define manual mitigation rules to block L7 DDoS attacks.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-77757603833852a766117782bce9f465a293d63b7355837371e62c3550241caa"></a>

## Direct properties — ddos_mitigation_rules / 4a7965fb5ffe / 3

- [block](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-548e7f44253f87015b78e2de1bd205a0508c2b471dd2dc9717fc7f09fd07bb24): complete subsection reference.

- [ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875): complete subsection reference.

<a id="canonical-54c6f984aeed4b61b4db94b9234b4f06c6fc230575c2012719b2fa3e647556d5"></a>

<a id="canonical-916d92f240e8f456eb850c7234aa20fe7b19224abd73e296a24d8f975678af7c"></a>

## expiration_timestamp property — ddos_mitigation_rules / 4a7965fb5ffe / 4

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4b746b55abfe51532c78eef1854c4c03f05965272449fb6443e6041be28716dc): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-90a2e84e284d41942431c422cd0d5dcc1554558590684c24bdb35d313765f975): complete subsection reference.

<a id="canonical-eaf02e8293d20668ac9ae00323b6f2eed1b8278117bf67c6d1daac98e25d569e"></a>

## Next pages — ddos_mitigation_rules / 4a7965fb5ffe / 5

- [ddos_mitigation_rules.block](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-548e7f44253f87015b78e2de1bd205a0508c2b471dd2dc9717fc7f09fd07bb24)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875)
- [ddos_mitigation_rules.ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4b746b55abfe51532c78eef1854c4c03f05965272449fb6443e6041be28716dc)
- [ddos_mitigation_rules.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-90a2e84e284d41942431c422cd0d5dcc1554558590684c24bdb35d313765f975)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-548e7f44253f87015b78e2de1bd205a0508c2b471dd2dc9717fc7f09fd07bb24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-041bea22bac1afb931119feee8ac40cd673354a54d9db7ac95093209078874b7"></a>

## ddos_mitigation_rules.block — ddos_mitigation_rules.block / 963735316d82 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- ddos_mitigation_rules.block

<a id="canonical-540c241037dd7776cdfb307600dbbb9464537a91d4ea1cfaa56cf20e2e51191b"></a>

Type: `"single"`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-d854dd7f368cb829006a72d116f6ad369bfe1d29ead01103247742f3d244cc02"></a>

## Direct properties — ddos_mitigation_rules.block / 963735316d82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1f88fa2ed8302388549d5e98125c7b5433e969f4c85682ba6228eb289b51d84"></a>

## Next pages — ddos_mitigation_rules.block / 963735316d82 / 4

- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d71c466a953d9c9177eb3fb7f65c75a1b9c3c0d357a998adce751f348ca1e003"></a>

## ddos_mitigation_rules.ddos_client_source — ddos_mitigation_rules.ddos_client_source / 8ddf6d7b5848 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-c749bd8f6eb26a75f2da7ca3790bbfe6582abc26579f946c0342b7d8d6f529ed"></a>

Type: `"single"`. Computed.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Upstream description:

DDoS Mitigation sources to be blocked.

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

<a id="canonical-ef6d451966cc1f71ced605ace4dcc4be2e91a1317530872066035ba08accaeb1"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source / 8ddf6d7b5848 / 3

- [asn_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-367c932d44d8ae68517e1aaa2ccc094da160cae2b3d287baede5f153a0a649c3): complete subsection reference.

<a id="canonical-710b6d8802545b415fb2cc8a88ede4ce73ee196334781f39e14e86ab54a416f8"></a>

<a id="canonical-b1406ee4f38e74b54be9c5f2eeecbcb01ab595e3be005bc97d6e64523ad91a75"></a>

## country_list property — ddos_mitigation_rules.ddos_client_source / 8ddf6d7b5848 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Sources that are located in one of the countries in the given list.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4423643394f2c676b2fe2e57c33864dd852321f7b974adc209d41460da3a1d64): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-40bd311ff4898f5a67353c113219076e8025e4516f39ae20c78bb056f1a1c720): complete subsection reference.

<a id="canonical-66fb81fc1f261cdf4557f17936f192d7a2bf9de4aa031250a8dea45c86fbb033"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source / 8ddf6d7b5848 / 5

- [ddos_mitigation_rules.ddos_client_source.asn_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-367c932d44d8ae68517e1aaa2ccc094da160cae2b3d287baede5f153a0a649c3)
- [ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4423643394f2c676b2fe2e57c33864dd852321f7b974adc209d41460da3a1d64)
- [ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-40bd311ff4898f5a67353c113219076e8025e4516f39ae20c78bb056f1a1c720)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-367c932d44d8ae68517e1aaa2ccc094da160cae2b3d287baede5f153a0a649c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-453e786528b95c4045b4638ce517facbaede5dc2dddc56d673b033eb62c03802"></a>

## ddos_mitigation_rules.ddos_client_source.asn_list — ddos_mitigation_rules.ddos_client_source.asn_list / 6f706e77a741 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-590fdc9b48552de5814bb7c165dc5bb0baace860da15390fb7c4828ac6b2b634"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-9363588aea306c2694fbe32af2960bf761a5e7950be65ce16109f9ccca2ef747"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.asn_list / 6f706e77a741 / 3

<a id="canonical-e96d5594482c42885454b675a547b0d86b4bff745f474feafcd03348be7f5791"></a>

<a id="canonical-07e29b18e772fa59607f544c5bf7a54e839248b00eafa6d14b311115a778bdf8"></a>

## as_numbers property — ddos_mitigation_rules.ddos_client_source.asn_list / 6f706e77a741 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-145c230fd805186da919ec394514e9f755f858820d918693ab5ffcf73b9ae277"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.asn_list / 6f706e77a741 / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4423643394f2c676b2fe2e57c33864dd852321f7b974adc209d41460da3a1d64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebb96baa42e1e34e005ab2a88aadf8ee749493499c32ae185f6e1d2ecce3d0fe"></a>

## ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 238eb003abad / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-c561a0038afaf8f9c5011383655429d224a216f2615d992e6c18d4c810b7be88"></a>

Type: `"single"`. Computed.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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

<a id="canonical-718eb348fa784f7358e2baf50a2ca01e8b53a986b92c702898a6393dec6f72c7"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 238eb003abad / 3

<a id="canonical-cdc214f9d9c3e4eaaf5a07df44090891bf1a9e6462cf927df924d12724125b93"></a>

<a id="canonical-78fbf06b63a807d2e45cf92b50a574bd687fc418b4d19baec0a1f29afe4213ef"></a>

## exact_values property — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 238eb003abad / 4

Type: `["list", "string"]`. Computed.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1d685bb11d81fc6c1813c99bd2609e8ce38179c6c87fd2b8548c5203ea30f80e"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 238eb003abad / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-40bd311ff4898f5a67353c113219076e8025e4516f39ae20c78bb056f1a1c720"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b68befbecf6ee6f5a26cfd536169a5e5bc7afeb923af49db1ceea776ea5c283a"></a>

## ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / bb06d87efcef / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-9fc345ac7b98ea4692c4bc095337abf0a8be41d989c7d7185ed57ef310789b5a"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-e746b0b1ce18a823f5e59bf10bd035774e4c38405a359ed3f41ce2a493e34ee5"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / bb06d87efcef / 3

<a id="canonical-f848fb3d1d116f16cf0d65824eeff619e121e9d5da94fa6c93db7ad812ac598e"></a>

<a id="canonical-72b502a44d3b19560e4e9246c0fa0969f168a4e685107300c3fb6ee06fcb6672"></a>

## classes property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / bb06d87efcef / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2cb45bce9d1ec1ef78b53dac00c6c9fc5d2e2162e7cafe1966cea01e9a84c29f"></a>

<a id="canonical-ad3ba3bc5a062e8abce2e0d0e81a469b00c22d4f8e24d72e9b015d78f7074bb5"></a>

## exact_values property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / bb06d87efcef / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9c76c65828ecc742a743b245363241a32699e365e21a5a40250bcf7e3b71757c"></a>

<a id="canonical-ff870aee9f5812b64ee08a90c71debf73dcf56768c200053a4c554ea6af910aa"></a>

## excluded_values property — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / bb06d87efcef / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f4671d943737fa8484c6005e5987f4879b0e6d3e36026e8546717c6decbae324"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher / bb06d87efcef / 7

- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2dead5072304fd7d4f4e59c8f4e8f0cbb01fbcce4d7fba429ed3ee8961861875)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4b746b55abfe51532c78eef1854c4c03f05965272449fb6443e6041be28716dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a854cf7ee40206a78c6e3623f7c8d41d413389d6bcb089794c5d7eeea1233c1"></a>

## ddos_mitigation_rules.ip_prefix_list — ddos_mitigation_rules.ip_prefix_list / 7d436d0e258b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-7d1612c64b6cb8efae5859701453f577c1481b696b9df642a4f607cca6f9199d"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-752ff8da30a01c9db22b7c9565437d14aa79c7d406a9e5530157afbe2bd2565a"></a>

## Direct properties — ddos_mitigation_rules.ip_prefix_list / 7d436d0e258b / 3

<a id="canonical-a3c32b08c425729f8b18175e99d725d527376c251a53e11079d736413361a212"></a>

<a id="canonical-7e6929b9cde8caeb870498dc4c81d34b2651d6c7e81f21e1f46f7343d27f0566"></a>

## invert_match property — ddos_mitigation_rules.ip_prefix_list / 7d436d0e258b / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-ee72e52a15d50ca57bc4e80743f4624e2ba0829178575925bb08b7d418a9fd01"></a>

<a id="canonical-ddbcac97778beb0e5e95419d059bf1498063f49cb4884277886d2334e7073d8c"></a>

## ip_prefixes property — ddos_mitigation_rules.ip_prefix_list / 7d436d0e258b / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ab15210a47468f097e0b4cb3b521469ce001f7b31a3979e6d26df8a9f4fef38d"></a>

## Next pages — ddos_mitigation_rules.ip_prefix_list / 7d436d0e258b / 6

- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-90a2e84e284d41942431c422cd0d5dcc1554558590684c24bdb35d313765f975"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e5fd5737111507276dba5810a542b4206f2c9a71167425e3545e5d8238f086c"></a>

## ddos_mitigation_rules.metadata — ddos_mitigation_rules.metadata / 8fede66cc527 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- ddos_mitigation_rules.metadata

<a id="canonical-df5c05a3ec9012240515e6f8383886547f56ba18f819e56bedb930ee75f553df"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-8e0dd755971e99b12c3d0836106ba491b1c37a2d548bbbafdf04c7d1bac39f6c"></a>

## Direct properties — ddos_mitigation_rules.metadata / 8fede66cc527 / 3

<a id="canonical-6e24f29e230c14f34c63d022501e8aaf1cf05103f207bb15384bbd33608de496"></a>

<a id="canonical-7c120fd53e7a70a978c62455a163f87bb4acc432f9834436c99b7d36dcf55922"></a>

## description_spec property — ddos_mitigation_rules.metadata / 8fede66cc527 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-32c729fe38bbe7dae5d5e306b5eb9d29376dabd36c10a963979ecb04875b4c6b"></a>

<a id="canonical-3c0636b8ebd6a23d66ad3f81b4fd2a793178f80c20d754f2d3f900f699546e6a"></a>

## name property — ddos_mitigation_rules.metadata / 8fede66cc527 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-8d5dadf937932ec9217847d52ed0f2d87872de2273fda5ac8aabaebb67a0a233"></a>

## Next pages — ddos_mitigation_rules.metadata / 8fede66cc527 / 6

- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-23fb934fdc8c11aec4fa330f9fdabf8e39499a5f76a8b4537af5d91bb4c53edc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7e7f4c028619bb61b3a22fba07eaf909a5ea14f5fa6a39e44d8d2aec7140a9c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c60d132492443dab0aaf36b3269c0fa6dd4ba51fe040ee3a602c20158aaf7e34"></a>

## default_cache_action — default_cache_action / 21db0c58f4d9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- default_cache_action

<a id="canonical-8c472aba7c7a70a85d8b66c9ba088020a56d417d514b2d6b066abeede87129f0"></a>

Type: `"single"`. Computed.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

<a id="canonical-2055ffcac34f00114393368ad47c07de7880497be3a79818d6578cb36497c982"></a>

## Direct properties — default_cache_action / 21db0c58f4d9 / 3

- [cache_disabled](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1947917c62477644b9b2cb51596024dbc16d7a814989ac5d768fd1e3f42d139e): complete subsection reference.

<a id="canonical-ab846aa680e4fe1bf91618e88df28987a557ef5ff8a4d363c423874c0ed1be51"></a>

<a id="canonical-79c29066f6b3874e9163eb7f14ea36587b7696a4745d88ab8bc7c4b225fa3fee"></a>

## cache_ttl_default property — default_cache_action / 21db0c58f4d9 / 4

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-a1b15159302dfe1eab9d4e23af7e2fb2a001c5a94e57956836c3706b3dc5f3e6"></a>

<a id="canonical-91eb9d878b744c954eee18ccb99826c16151c18647bcfe81db3dec968906dd95"></a>

## cache_ttl_override property — default_cache_action / 21db0c58f4d9 / 5

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-293f30d0c47ed2fc3de6f1d06f451bdb70ea7747563f9af29e44f2783301f787"></a>

## Next pages — default_cache_action / 21db0c58f4d9 / 6

- [default_cache_action.cache_disabled](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1947917c62477644b9b2cb51596024dbc16d7a814989ac5d768fd1e3f42d139e)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1947917c62477644b9b2cb51596024dbc16d7a814989ac5d768fd1e3f42d139e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1848e52aa51a6c12956fdaf83aba206a14691aef19bc84ec70851d264d6f798"></a>

## default_cache_action.cache_disabled — default_cache_action.cache_disabled / 6d1a1f244d89 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [default_cache_action](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-7e7f4c028619bb61b3a22fba07eaf909a5ea14f5fa6a39e44d8d2aec7140a9c7)
- default_cache_action.cache_disabled

<a id="canonical-9d198f19d2879e1757ae694c2fac1334d7158da7ebca1c043b3b5c415599cd1c"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-a4a8c9a98d40bf8746ffdd8a4579825fea976eb3d0c4ef2cb5892ed629947a10"></a>

## Direct properties — default_cache_action.cache_disabled / 6d1a1f244d89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-70c1e8d78cfb72aea4d5382e7f402ec4b6ec001c36c2b5ea6ddac0c425cd42d9"></a>

## Next pages — default_cache_action.cache_disabled / 6d1a1f244d89 / 4

- [default_cache_action](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-7e7f4c028619bb61b3a22fba07eaf909a5ea14f5fa6a39e44d8d2aec7140a9c7)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-44182d76825e878f69f595ae497de7899cf9189729eada12d24e7570c7aceadc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab44d419f34ec41c977565ae3a2b3a8195adf5c179801b8e1b16fc1b3aab8cc0"></a>

## default_sensitive_data_policy — default_sensitive_data_policy / e3e51eb44725 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- default_sensitive_data_policy

<a id="canonical-bc7c19993cbc145c3b22eb3c3ce3c5750837a65e4df828d6ad42e7d859dc1113"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [default_sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-bc7c19993cbc145c3b22eb3c3ce3c5750837a65e4df828d6ad42e7d859dc1113)
- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-9abebbf10799beafcd7cf8ee9287f1ff4c96552d46c31c76ef8d1aa266f6081a)

Select alternatives according to the provider validators above.

<a id="canonical-eba997164a5e57909d2b1a62d8500ff66e3cebf88a6ef0be6c9292307d72ce87"></a>

## Direct properties — default_sensitive_data_policy / e3e51eb44725 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-52664bf407423f37a95fbb4a62c5796c0fa3c6381c0f0d92a1e15953252a7961"></a>

## Next pages — default_sensitive_data_policy / e3e51eb44725 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-df66c2908ab3c0ecf93a8cd22af687b9430c25b9971c0b8a9e6ba728a2001a33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0fe009f5d4c9fbe49c78fa0c7a971a99957a52825bbb3d517cf534c657110dd"></a>

## disable_api_definition — disable_api_definition / defcef4b6f65 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_api_definition

<a id="canonical-1b6e410e8050a4c9750180748a209fbcf3b07cc1723c1f135fed669c8a85c585"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-7d5582ece2fee89e6e6b17b4e4e55affb0e8587ffbe8cc2d6163818c375219d0"></a>

## Direct properties — disable_api_definition / defcef4b6f65 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a569b318c9d78d6e65569e064acd3b64ecdcd51beac9232b5f695d3da413d0b6"></a>

## Next pages — disable_api_definition / defcef4b6f65 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a4d325e03dcc681fa8b89708c9bdb4988252e05b85d3696de9e0565db4361b9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cf7c9edceb0326fc36b10c220d517de867bdd625c71559a354ae1fffade4fb6"></a>

## disable_api_discovery — disable_api_discovery / 49695cb37f0e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_api_discovery

<a id="canonical-a102b918c4995cc60d8880c031ee12cb9699da98da66742f4f69607a43d8a6d1"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [disable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a102b918c4995cc60d8880c031ee12cb9699da98da66742f4f69607a43d8a6d1)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-c032068c0b2ef3b0b272d856d27da5939179b058f74d97ce68a1b7345dd75b5e)

Select alternatives according to the provider validators above.

<a id="canonical-49181bf3a10fdaccaada46da8517b60b95430b70c82887e0527a688d09d74f34"></a>

## Direct properties — disable_api_discovery / 49695cb37f0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1de360435106d7e8c3b6b950cc3eb092ec842b9eb4e7e04b686bcc08bba522b"></a>

## Next pages — disable_api_discovery / 49695cb37f0e / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f470ae3f8b5dcb8ac16e8a1165c62a28e5a7546a1d389d547cc001325e5a2b47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13b59027b13b8163e8725cd88592e9b1e6c3290f98c4dc3f3bb7aedbab9091f5"></a>

## disable_client_side_defense — disable_client_side_defense / 64adc25ef427 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_client_side_defense

<a id="canonical-a885b2a10be231e4132d2d07c964ee4f3a3b106473d40194f7fbfd22fbd77c04"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3238465d53a61c8feb606a0187629839c747185c3e5d2f65e00470b12a8075d7"></a>

## Direct properties — disable_client_side_defense / 64adc25ef427 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b8e6d8b6b47dd33ff43c002e5d8f7f1e4901e5161c609301626224b93f2db098"></a>

## Next pages — disable_client_side_defense / 64adc25ef427 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6cf2ba0eb819fccf9e4e9e9f5993fa8e2dc5fafdb640cb4ddf1a83840aab34d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77968d70b494a619a3e3af95b48a23d0cf386c6e9fb2f03483a4c251ec46421a"></a>

## disable_ip_reputation — disable_ip_reputation / 1ecc86f33def / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_ip_reputation

<a id="canonical-1bf7b5a83d4e57d9e96e810aaccc6dcafe95f59c58c956933c1d60bdc0e252c4"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1bf7b5a83d4e57d9e96e810aaccc6dcafe95f59c58c956933c1d60bdc0e252c4)
- [enable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-06bc599ba9a0a22102b4b7d3972f4defd6f1b03c273193df0423bf48581b56a0)

Select alternatives according to the provider validators above.

<a id="canonical-d17076983882d4262c5835358b129f81ee76f325a53e45eb843699a000504f97"></a>

## Direct properties — disable_ip_reputation / 1ecc86f33def / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c10ea885d185bd15f2d057f49de768acf2d82a06254a45a2b5d163a24d2e4f7a"></a>

## Next pages — disable_ip_reputation / 1ecc86f33def / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-606be885391d214dacd98bf078dd9ed9168af7b9d43bf8d452c8c2ed54e1002d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad47d8d45586ddbb0abfcb239584253ee84f441222dc8a6d44c9e6dd5b170753"></a>

## disable_malicious_user_detection — disable_malicious_user_detection / f4dd4d4c3ac8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_malicious_user_detection

<a id="canonical-e0af114a9e3479389ac7311c897fbe0e352cf13ebf5ec76f59280678f35e9384"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-e0af114a9e3479389ac7311c897fbe0e352cf13ebf5ec76f59280678f35e9384)
- [enable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3220cb9e442f7fa713187eebb8c8d943e1a28aba7d54521d79cbd7c47a3b7757)

Select alternatives according to the provider validators above.

<a id="canonical-d1f767ea357f8b32bef12ef7bc98485c1c77ed418cf22a4c21954a33ec8342f1"></a>

## Direct properties — disable_malicious_user_detection / f4dd4d4c3ac8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9b4759ef461cb252381f9e472f376773c9ede58187cea6f4de63914e541086cb"></a>

## Next pages — disable_malicious_user_detection / f4dd4d4c3ac8 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-93bf6479464cd8a577214ea9fd9a2749eb0acbcf622374fdd47b73cd19a4d798"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a943388011270474eb67aa18074a0cbfb55038beb47b8fcd7f0a240cab45110d"></a>

## disable_rate_limit — disable_rate_limit / aeb1f2e449ae / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_rate_limit

<a id="canonical-2c81ec2c0fd31384f5c76c0244636639b30ae4efecd0932372c1f1ed12ef0164"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable rate limit.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-8afffa23e49f26e1383966564f2350b60b1aa428dcbdc62fa553c4ec2d17340b"></a>

## Direct properties — disable_rate_limit / aeb1f2e449ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-088f64fe5590f0389d0ff40067fd89fbb469fae92b2286bac588e5cfb814d2a7"></a>

## Next pages — disable_rate_limit / aeb1f2e449ae / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-bc358ac1840523be4841d320886d9879cc831b43c07b4c9741eb51161800e999"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-953facabd71ac21a8282a99ef1f088464a1ad0b2b0e02b0db3ae5a0379c2b79c"></a>

## disable_threat_mesh — disable_threat_mesh / 4a059e04ada4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_threat_mesh

<a id="canonical-2eb79d94c6ca2f3fc3fef56cd37e84e6a139ffb5563f8fd5fe0144871b26ebb7"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2eb79d94c6ca2f3fc3fef56cd37e84e6a139ffb5563f8fd5fe0144871b26ebb7)
- [enable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-417e345f3eb73b352e523c3aa18328b3c0bb6784884972f05b2fa28939cca1c6)

Select alternatives according to the provider validators above.

<a id="canonical-5e10c84dc95fac46d7f7026be42b876140f71033610ff22f1ce9a7fe0bb7a362"></a>

## Direct properties — disable_threat_mesh / 4a059e04ada4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bba94a38a39527581d2f0070dbf8f3b26ecdb6c9cd5d5f2ce0a6203de1726df3"></a>

## Next pages — disable_threat_mesh / 4a059e04ada4 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-45be9a9d0453fcbf1d5b7a9e640b3ddb6c47e33543d135435339fa8f45dc2a25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84d3defb4e47b3cd48f67747fd3be3512571d273963a1844a04daf3520928527"></a>

## disable_waf — disable_waf / 4f2362d86024 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- disable_waf

<a id="canonical-9a9191721466dc4820007fe1bf0e439cdf21d7b202a10265565cec3a2293921b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3417cc7c690991a650b1236043752c93a4229266291559624a8d89066a1611df"></a>

## Direct properties — disable_waf / 4f2362d86024 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c3688e8b34170ea89eeb3e8f628a783c17f4c4d475b292ca56510c60d956589"></a>

## Next pages — disable_waf / 4f2362d86024 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a1999199a4bbda79a78fb124abf77332b28b07ab86103ffa79310bdddba181e"></a>

## enable_api_discovery — enable_api_discovery / 2956d51d7ac8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- enable_api_discovery

<a id="canonical-c032068c0b2ef3b0b272d856d27da5939179b058f74d97ce68a1b7345dd75b5e"></a>

Type: `"single"`. Computed.

Specifies the settings used for API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

<a id="canonical-5a4f9860effa68abaeb65a809bab4cae0204dbc072576389ba920d7005ad6773"></a>

## Direct properties — enable_api_discovery / 2956d51d7ac8 / 3

- [api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e): complete subsection reference.

- [api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0): complete subsection reference.

- [custom_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d8cea2c986d75f0f461fc48f5299fafbc7f28ef79c4c8fd12b7a5863cda65736): complete subsection reference.

- [default_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-924655935a29713b71d9683ada5949479dc9e8a942b2280247cfdf4c3e580b83): complete subsection reference.

- [disable_learn_from_redirect_traffic](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d0ed52e1a86221b80749c741711d70dbebdb00aa15bccae4614bd04e4df057a5): complete subsection reference.

- [discovered_api_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b22017f4a2fc2d6e6311719ed49308077d507d37fc48b53cd22e86f5c24ba279): complete subsection reference.

- [enable_learn_from_redirect_traffic](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-7c48c0b9fc1142de5fa4f1a82551b95191251e7c305bcd8950ef84b0d4709770): complete subsection reference.

<a id="canonical-fa10c72d50e39fdd4209f11ee8118fc27e2e01c2ddca343842a2b3955298ebf5"></a>

## Next pages — enable_api_discovery / 2956d51d7ac8 / 4

- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0)
- [enable_api_discovery.custom_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d8cea2c986d75f0f461fc48f5299fafbc7f28ef79c4c8fd12b7a5863cda65736)
- [enable_api_discovery.default_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-924655935a29713b71d9683ada5949479dc9e8a942b2280247cfdf4c3e580b83)
- [enable_api_discovery.disable_learn_from_redirect_traffic](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d0ed52e1a86221b80749c741711d70dbebdb00aa15bccae4614bd04e4df057a5)
- [enable_api_discovery.discovered_api_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b22017f4a2fc2d6e6311719ed49308077d507d37fc48b53cd22e86f5c24ba279)
- [enable_api_discovery.enable_learn_from_redirect_traffic](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-7c48c0b9fc1142de5fa4f1a82551b95191251e7c305bcd8950ef84b0d4709770)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99f4b2dacffbc21c4798c6413c075c98905ae8873a2324f4f9c769d1ef566bd8"></a>

## enable_api_discovery.api_crawler — enable_api_discovery.api_crawler / 0f053c524ca3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- enable_api_discovery.api_crawler

<a id="canonical-4ef1c214ae1b299a2c6fb6961934d87c9b79619a39d65d15f418677fdda23b7e"></a>

Type: `"single"`. Computed.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

<a id="canonical-f68f040734e81c092e91197c8a729ee05c8106dd1fba9b5cbe8b3566f02c897c"></a>

## Direct properties — enable_api_discovery.api_crawler / 0f053c524ca3 / 3

- [api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee): complete subsection reference.

- [disable_api_crawler](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1cc475c26c933a448b7721dc3dd8eaffd58e151d50676ea5562d8bdce2a17fe7): complete subsection reference.

<a id="canonical-b323bada73f3c37801adba3b2a68c5471f4074b504ea4bc3d9de184c1d421d3c"></a>

## Next pages — enable_api_discovery.api_crawler / 0f053c524ca3 / 4

- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee)
- [enable_api_discovery.api_crawler.disable_api_crawler](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1cc475c26c933a448b7721dc3dd8eaffd58e151d50676ea5562d8bdce2a17fe7)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76b6cea724741820c71610544a551dbd19319a666a8be2d4764bd2da14e2cbfe"></a>

## enable_api_discovery.api_crawler.api_crawler_config — enable_api_discovery.api_crawler.api_crawler_config / 21eee6497677 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-5901c33394be96bd8d72c9862e0ea3a9ce7f4977725b19d8560a07134af113f8"></a>

Type: `"single"`. Computed.

Crawler Configure.

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

<a id="canonical-b9c432621c39b10472e47a39f7f24c8f63c0eec92380638d45b13f45a578250b"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config / 21eee6497677 / 3

- [domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249): complete subsection reference.

<a id="canonical-04ac4f7906abc9a777bec2325f34e1b5de490fb8f7478041894201778a7f7b6d"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config / 21eee6497677 / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a133320956d9f95bc5108a619a47801771a74a0e241f72394f1ad62e75ed81b7"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains — enable_api_discovery.api_crawler.api_crawler_config.domains / 7952e14e38dc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-0128ed7fc9ad9dd1eb61141a5213c9c80b1aa48cc5d43a5737d706efc5ff6534"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-50018cc5847e5e43c0929473bd6f20d9e52299bdfd28df5d685809a4edb81e8c"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains / 7952e14e38dc / 3

<a id="canonical-2f865b10789f669747923957ad6d6ee6de79cc3a9fafa02b19a2012a6c7be263"></a>

<a id="canonical-83f9794dc378ea767beb6fdebd2c58572bbe78bef4c19fbc3a7c0b09e06be982"></a>

## domain property — enable_api_discovery.api_crawler.api_crawler_config.domains / 7952e14e38dc / 4

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-50f72f4773ff6a6ce1a50b2b5f5d31ea4ecc61547d908e2172a9a4be76e7b9ad): complete subsection reference.
