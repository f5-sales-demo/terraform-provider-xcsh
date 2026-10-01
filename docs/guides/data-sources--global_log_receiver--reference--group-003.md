---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-f110068ecff5526412e84df55b8269c1fdd92791435ab9eccf142fa854c650c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f7f9493a7213faea0163447b51eb338fa74da92a9848971130b46320b73a17b"></a>

## http_receiver.batch.timeout_seconds_default — http_receiver.batch.timeout_seconds_default / 2c9497d40962 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1)
- http_receiver.batch.timeout_seconds_default

<a id="canonical-0e9d0c4960e1c8c66948d26deb3076da91b971d6ae0ebd37aa364c4ed017d487"></a>

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

<a id="canonical-29dc5212fa24a2d057f61120509ca1e19f0a986ce51b98fc91e87e1bbfa83534"></a>

## Direct properties — http_receiver.batch.timeout_seconds_default / 2c9497d40962 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58926592d92cc3f054cb296e77116d77db365fc05967d071e6b67d5e43d5e58e"></a>

## Next pages — http_receiver.batch.timeout_seconds_default / 2c9497d40962 / 4

- [http_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-7b3f421e3e1e620d267fe8dd529baf931f882281796895b2036ac4525303aae1)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3a9ee8e728eb72b481a05db9cfb462e0fb5a753a9e91f1f13c118838787c514"></a>

## http_receiver.compression — http_receiver.compression / d33d51dc41bf / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- http_receiver.compression

<a id="canonical-6c50d54b945021f0bdb741480c9cebdc1965f378f118522a51b2784abe7b97db"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-0306e70b5fef2d5cccb5eb0e579a21b56469afde8c7cbf9102451592cb3a9c95"></a>

## Direct properties — http_receiver.compression / d33d51dc41bf / 3

- [compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-3b3eb6b88bd0381ec54ef681ecda979df8d5664687ab4280fb11723f057fd186): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-a034982347e49026bade97405ef91c6b69211267a3d5ca5535203925c5dd4cf9): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-f32ece9826f88c11ba6a4b81e63b39a0634e0b021f78e747f211505100e150c2): complete subsection reference.

<a id="canonical-23ada13b3f5f30c9415b53fa9645b26c681ca11e404d3ec93f6535aa7e939d4d"></a>

## Next pages — http_receiver.compression / d33d51dc41bf / 4

- [http_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-3b3eb6b88bd0381ec54ef681ecda979df8d5664687ab4280fb11723f057fd186)
- [http_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-a034982347e49026bade97405ef91c6b69211267a3d5ca5535203925c5dd4cf9)
- [http_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-f32ece9826f88c11ba6a4b81e63b39a0634e0b021f78e747f211505100e150c2)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-3b3eb6b88bd0381ec54ef681ecda979df8d5664687ab4280fb11723f057fd186"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35409ae786ce245b5957c9e52ea1015cd8a1b14dbfe27cd5d7b5ad0726a95c67"></a>

## http_receiver.compression.compression_default — http_receiver.compression.compression_default / 9e574dc0790c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962)
- http_receiver.compression.compression_default

<a id="canonical-5065dd38ee2bbe8cacfe882c751b30859186541dd69f7bbcd04f19516accc6ba"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-fc351f03eaeb9a517a3ad0945b9295823a31c25c9d7a4d9896cc7c104b98191c"></a>

## Direct properties — http_receiver.compression.compression_default / 9e574dc0790c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-776d707cb9db1a6ae221a21f199a218cfe049dd07884a773cd55cd0209f9dde4"></a>

## Next pages — http_receiver.compression.compression_default / 9e574dc0790c / 4

- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a034982347e49026bade97405ef91c6b69211267a3d5ca5535203925c5dd4cf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad574dfdcbfd88a0f3e6009e196752f500acf7fd6534d2e9cd9bc5b170e6bcd4"></a>

## http_receiver.compression.compression_gzip — http_receiver.compression.compression_gzip / 39273f221627 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962)
- http_receiver.compression.compression_gzip

<a id="canonical-b181489196ef22dd3bc1bddaac4aa20eeeed67a40af1b028fdf931ecd031476c"></a>

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

<a id="canonical-168968860852663e301cdac9b6fb428a15c6d1c08c53efbdf327b449430b2f28"></a>

## Direct properties — http_receiver.compression.compression_gzip / 39273f221627 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f2581be4ec56a497b5986a6da1debd4583f40db0a1f55a3cd93ed19f14e2e9a"></a>

## Next pages — http_receiver.compression.compression_gzip / 39273f221627 / 4

- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f32ece9826f88c11ba6a4b81e63b39a0634e0b021f78e747f211505100e150c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa59662b520726e0aba096bd4699523ec9c7f39de0eefbe68755590edd87dc8e"></a>

## http_receiver.compression.compression_none — http_receiver.compression.compression_none / 09f50c9a1b49 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962)
- http_receiver.compression.compression_none

<a id="canonical-f92af2ac5d27141a3a0392c3bc013640a3bef8245ea45bae976d625ea43d1ee3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-8bf2a54a8e4a5995bb32819e931bc431818570db8baa5679fe0cb928aec7b46c"></a>

## Direct properties — http_receiver.compression.compression_none / 09f50c9a1b49 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4000314fe219e519688e91ccb48d5e77d3085b0c84bff644abbf036d2502b1f"></a>

## Next pages — http_receiver.compression.compression_none / 09f50c9a1b49 / 4

- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-5dae3adc806a3cca4b0ba9fa98f97475c3e566148ab3782f747407bd10fe1962)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e3aaf3b727494a85b078661c9bb0be0422abd4b2ac37efada8cbf6b9841ccf2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5106fd98e0a496ce3e6cf59c112ef7cb832c2d70dda1fa24d0f5c50d233f53bf"></a>

## http_receiver.no_tls — http_receiver.no_tls / 771967b96bad / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- http_receiver.no_tls

<a id="canonical-7572b41242e440e6f6f4bb5cdf52118e1e78c2dc2a3148fe81fe679ee30794b8"></a>

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

<a id="canonical-df3cae01209f3270655996c8f3184d5f11fc1fc51ce037d9fe9b67fb5f1681b6"></a>

## Direct properties — http_receiver.no_tls / 771967b96bad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b0661fd0d471010b6fb95d2932957ebf03bca6824e099a40ef4b2e1402313fa"></a>

## Next pages — http_receiver.no_tls / 771967b96bad / 4

- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d30cbc26b6e8fc138e463a92fd6474442b0ec61f27dc057d0c898a8a92251a4f"></a>

## http_receiver.use_tls — http_receiver.use_tls / 482a3a151164 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- http_receiver.use_tls

<a id="canonical-aa4e8b8d8d38465ca5a6fd8f3ec7dc36981977b167cf2705a2bd00c48acfcf98"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

<a id="canonical-77375a50e99caa6f93532ac72b388dbc499bea868d7f4f558eb165d56136638f"></a>

## Direct properties — http_receiver.use_tls / 482a3a151164 / 3

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-d9905f2b055405114c76619e922dd5aca8488ad1bab69dd0ee18c04d7e752a04): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-0bc58580d86bd67d5ab2a7ada935b10536167dbf0a9f663da5882a029eb20391): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-18725a8ae19a9413518a4581655205f013e4178a9db24df2f064deb05cf077df): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-15e55f6158578755599c94e8ea4b842849be94b2fd8102834ca13fe0e9d16c8f): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-699e678e8cd8c6649117f06b8750560d26e90f76334db9b4b55b6338b3ad4bfc): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c7046f90756d40ce2c5b7ae38eb8a1b2db5adece934681927b74ac6305800ace): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-d797009afd55eb5d22cc3b0f10c9e9c0b281fa6334071eae84773dd05f844f2b): complete subsection reference.

<a id="canonical-32df6a960ccdf166c9a071f9db1ddbb3e5bcb94e8c9952961f58c8180340e6ce"></a>

<a id="canonical-cc314ae440545fcedafcad18592dc246d070e60d20c960218b650393e9a6efb8"></a>

## trusted_ca_url property — http_receiver.use_tls / 482a3a151164 / 4

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-9eedf6feb28e5c49cf151328ffe7a0862cd11982fc80d3f0e09277658c24f690"></a>

## Next pages — http_receiver.use_tls / 482a3a151164 / 5

- [http_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-d9905f2b055405114c76619e922dd5aca8488ad1bab69dd0ee18c04d7e752a04)
- [http_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-0bc58580d86bd67d5ab2a7ada935b10536167dbf0a9f663da5882a029eb20391)
- [http_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-18725a8ae19a9413518a4581655205f013e4178a9db24df2f064deb05cf077df)
- [http_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-15e55f6158578755599c94e8ea4b842849be94b2fd8102834ca13fe0e9d16c8f)
- [http_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-699e678e8cd8c6649117f06b8750560d26e90f76334db9b4b55b6338b3ad4bfc)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c7046f90756d40ce2c5b7ae38eb8a1b2db5adece934681927b74ac6305800ace)
- [http_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-d797009afd55eb5d22cc3b0f10c9e9c0b281fa6334071eae84773dd05f844f2b)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-d9905f2b055405114c76619e922dd5aca8488ad1bab69dd0ee18c04d7e752a04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5de7b41512c830f9fca9ae672aba9b89e7eb1c2d54a885952cab355d5a337656"></a>

## http_receiver.use_tls.disable_verify_certificate — http_receiver.use_tls.disable_verify_certificate / 062c891af332 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- http_receiver.use_tls.disable_verify_certificate

<a id="canonical-46d7ba665220528fbb56e00e55445f977dc42ee231f6810243f3a8cf4ace3be8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-3e0738efa3508bda48e1c81a5e3543bcd94f6c837992eee9296e5a1e813f767b"></a>

## Direct properties — http_receiver.use_tls.disable_verify_certificate / 062c891af332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1433295a204f840fd587e48be7756245fd2ee3a17dbbbbaa6e61107bf2141cf"></a>

## Next pages — http_receiver.use_tls.disable_verify_certificate / 062c891af332 / 4

- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-0bc58580d86bd67d5ab2a7ada935b10536167dbf0a9f663da5882a029eb20391"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e60d638d8853c7d55082691cd64d6066de2b5b325f21bfa710c6b8ed92fc1936"></a>

## http_receiver.use_tls.disable_verify_hostname — http_receiver.use_tls.disable_verify_hostname / 51f349677943 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- http_receiver.use_tls.disable_verify_hostname

<a id="canonical-b029dbb33f2ff2fcf14b77b8c3d35ef838dac463835453604660a5ab896a0610"></a>

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

<a id="canonical-c7e026aebe6901c390b9738de8e6b7061c311c7a77319858ec329b676c6df033"></a>

## Direct properties — http_receiver.use_tls.disable_verify_hostname / 51f349677943 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c134018044db02c72b9360a9902dd1c54d36279cf466f21ed814e89d50dfaa9d"></a>

## Next pages — http_receiver.use_tls.disable_verify_hostname / 51f349677943 / 4

- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-18725a8ae19a9413518a4581655205f013e4178a9db24df2f064deb05cf077df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44b5ed5dc97c87aae4de947b47ad1bac4c7557772d5cb69c19c8640a42c3a3d9"></a>

## http_receiver.use_tls.enable_verify_certificate — http_receiver.use_tls.enable_verify_certificate / 491eba3320f8 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- http_receiver.use_tls.enable_verify_certificate

<a id="canonical-eec642e6269ba513a0e83aee691e912859e583b5be6bb3177d372318be53029d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-4dea3d00671a198e2f082bf82a594f4bde24e70d547c75abcf3c8cf4edd3cd9c"></a>

## Direct properties — http_receiver.use_tls.enable_verify_certificate / 491eba3320f8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-793b11e0ba9889b4384885999e1f27834b008abda9c90d306161fe90eb690134"></a>

## Next pages — http_receiver.use_tls.enable_verify_certificate / 491eba3320f8 / 4

- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-15e55f6158578755599c94e8ea4b842849be94b2fd8102834ca13fe0e9d16c8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-534616f239987889510f68a716800b3e17691aa7dc4b91b132ed749d5aa2d5ae"></a>

## http_receiver.use_tls.enable_verify_hostname — http_receiver.use_tls.enable_verify_hostname / 865f315e6381 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- http_receiver.use_tls.enable_verify_hostname

<a id="canonical-f8e3737b10451e6c376af765f799c5d8c194f6b690634014f3e027fc05b4e22c"></a>

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

<a id="canonical-a6c8ba1cbc0e20caa0635d7a4a017a12d4ca8cef06b413eca364e77125cf161b"></a>

## Direct properties — http_receiver.use_tls.enable_verify_hostname / 865f315e6381 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dbc92ffb8003cb83c676b09579dff666e33bf0aeedd4dbb9d2ac0170c59a722f"></a>

## Next pages — http_receiver.use_tls.enable_verify_hostname / 865f315e6381 / 4

- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-699e678e8cd8c6649117f06b8750560d26e90f76334db9b4b55b6338b3ad4bfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1429df3b6a4f39ffa862c978bc5838ec58e681299a569f026450f0b82cf5bb16"></a>

## http_receiver.use_tls.mtls_disabled — http_receiver.use_tls.mtls_disabled / 258c96beea46 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- http_receiver.use_tls.mtls_disabled

<a id="canonical-04ab152c359632c0c484f1cd3db30fa367439c070883e6e49726351ef2bf3a86"></a>

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

<a id="canonical-39a891cc72f1ed9b0379921a76faec6f2e169e34266ce7da7eed27e4bfb9119a"></a>

## Direct properties — http_receiver.use_tls.mtls_disabled / 258c96beea46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2176afeca0a31073b64efabaa817bb4ad17fa3fc21cf8ccab64697a77647d5cd"></a>

## Next pages — http_receiver.use_tls.mtls_disabled / 258c96beea46 / 4

- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c7046f90756d40ce2c5b7ae38eb8a1b2db5adece934681927b74ac6305800ace"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-649e9bf9ee88e439070e0a4fe583f3018e505b4213624d9c54a7bed21a012ad3"></a>

## http_receiver.use_tls.mtls_enable — http_receiver.use_tls.mtls_enable / a5376ab1489a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- http_receiver.use_tls.mtls_enable

<a id="canonical-976329e75b15969ce6bfb460f4c83e19ebbbce612e89ff12091ba4f8b1d65ac8"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

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

<a id="canonical-3618fc70e7a80af8fceecce6369a0ed1e3832bf816e15f0b783c3cf02b3b98ac"></a>

## Direct properties — http_receiver.use_tls.mtls_enable / a5376ab1489a / 3

<a id="canonical-c5bbd499bcf3a08fa1fdd2a6db2eac07dba6f106f18879eee2eacaffe70bf9aa"></a>

<a id="canonical-52988358bd9691aa0b1f6714b1931b02d8b7d005bfd412dc322904470d2a72ea"></a>

## certificate property — http_receiver.use_tls.mtls_enable / a5376ab1489a / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-5f4c4c7108ba4239ccd10fd2d59649342d8b5989271092f8c4e8f7f7be4efe76): complete subsection reference.

<a id="canonical-519006e37ed34a02001c9573438f6b7421c45269ed59c3f8f225c3dfd205b3f1"></a>

## Next pages — http_receiver.use_tls.mtls_enable / a5376ab1489a / 5

- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-5f4c4c7108ba4239ccd10fd2d59649342d8b5989271092f8c4e8f7f7be4efe76)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-5f4c4c7108ba4239ccd10fd2d59649342d8b5989271092f8c4e8f7f7be4efe76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f7e3d582d62a1b363281cb56cffed98a6e80004243ab40f2f3ae621f87dcf1a"></a>

## http_receiver.use_tls.mtls_enable.key_url — http_receiver.use_tls.mtls_enable.key_url / 06ae10a4d5ac / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c7046f90756d40ce2c5b7ae38eb8a1b2db5adece934681927b74ac6305800ace)
- http_receiver.use_tls.mtls_enable.key_url

<a id="canonical-57caad8d342718d189f8cd6e6e71542a33060b56331b59f8e2986557e0593f12"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1cffe700c3be884f275998033425858ce15d6b79260b2359579f0b8cf6216362"></a>

## Direct properties — http_receiver.use_tls.mtls_enable.key_url / 06ae10a4d5ac / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-f52c8e4ac1bd1729b6fce82393a06e37ae3712cbda3ddbd7a123547b391c29b5): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-12db113b11d2f3b5e3160d457bbcb82da123b7b6e87edaa89bc1483644fc5ae6): complete subsection reference.

<a id="canonical-131124fb10aacd885b5df1e2a64c23cf09168a823202466073f4bda7fa6a3893"></a>

## Next pages — http_receiver.use_tls.mtls_enable.key_url / 06ae10a4d5ac / 4

- [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-f52c8e4ac1bd1729b6fce82393a06e37ae3712cbda3ddbd7a123547b391c29b5)
- [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-12db113b11d2f3b5e3160d457bbcb82da123b7b6e87edaa89bc1483644fc5ae6)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c7046f90756d40ce2c5b7ae38eb8a1b2db5adece934681927b74ac6305800ace)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f52c8e4ac1bd1729b6fce82393a06e37ae3712cbda3ddbd7a123547b391c29b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c81b3b13ecf142512797f26f70637a9426fa399edfc1646da2ce7841e97f44f"></a>

## http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 469a1269ad17 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c7046f90756d40ce2c5b7ae38eb8a1b2db5adece934681927b74ac6305800ace)
- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-5f4c4c7108ba4239ccd10fd2d59649342d8b5989271092f8c4e8f7f7be4efe76)
- http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-777904f0a020cc856bb741ef138c820d839ab3a0f73b560739800e334e089251"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-29cad8f98e68186c2bb2b04e26a74afc76ff1c852ee292e376b6a1d7684dfec3"></a>

## Direct properties — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 469a1269ad17 / 3

<a id="canonical-792ab75b7de294b95ac4e22ac67af053de3768cb4bcfac36b0af34ac1ff745a1"></a>

<a id="canonical-2883a3497385a66ba4c48728e0f393a9f7525033401353db4397fd1b2693a0d5"></a>

## decryption_provider property — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 469a1269ad17 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-c3c23170fc40dd16db9896bef93705e4cd57a33303c23bfc1f4cd16d9e843410"></a>

<a id="canonical-65d7f72c7a6249dd6519ccf04c1a6d07f41c6bfca0ff4e67fcd2dfdf83f15987"></a>

## location property — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 469a1269ad17 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0926218d8c16f367888bed8f9b45d88cdc6efa9fade3431b5ed3144fba057a5f"></a>

<a id="canonical-65498982e320164cd9fd5e0dbddff5a3566c1739282e0cdd309dee04f611a139"></a>

## store_provider property — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 469a1269ad17 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-b9b8ca39492f632f84973a822f3af5be612ac4430cdbebbf9608f8e982265059"></a>

## Next pages — http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 469a1269ad17 / 7

- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-5f4c4c7108ba4239ccd10fd2d59649342d8b5989271092f8c4e8f7f7be4efe76)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-12db113b11d2f3b5e3160d457bbcb82da123b7b6e87edaa89bc1483644fc5ae6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb98d44ef2cc47ac4623109b85c7eec1823f4beb45efd3bac02407c556c92cc9"></a>

## http_receiver.use_tls.mtls_enable.key_url.clear_secret_info — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 12afef4b87c9 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c7046f90756d40ce2c5b7ae38eb8a1b2db5adece934681927b74ac6305800ace)
- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-5f4c4c7108ba4239ccd10fd2d59649342d8b5989271092f8c4e8f7f7be4efe76)
- http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-e76a95671c4710ae4b95572185b18640078c4cf137759d7fdf5ac2c269f557e2"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-88567b44de8dca206522460c4bc55fb65cc8e2a3aa8c9b39aab480f245ac033b"></a>

## Direct properties — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 12afef4b87c9 / 3

<a id="canonical-466dbd2ac17fde3041f6a47258db106fa616d63eedd9296acb0c135c2ae08ffd"></a>

<a id="canonical-af2b8f36d9104847b8a33a2aeb57ea2a078b7fba15dde1848855af5aa684f234"></a>

## provider_ref property — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 12afef4b87c9 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-89656200f21bbb540241abaf7cebf4c2bbe9e962906f10d8c6037dcec302c0a6"></a>

<a id="canonical-00637e57078e0d2aaa26b595b79daf6d3e993fe0f15f047da20f177d89d2de5c"></a>

## url property — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 12afef4b87c9 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-bf9a09026b7a67e914da2823e04909b2b74991ca0f19e4df3bdafe57e7f813dd"></a>

## Next pages — http_receiver.use_tls.mtls_enable.key_url.clear_secret_info / 12afef4b87c9 / 6

- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-5f4c4c7108ba4239ccd10fd2d59649342d8b5989271092f8c4e8f7f7be4efe76)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-d797009afd55eb5d22cc3b0f10c9e9c0b281fa6334071eae84773dd05f844f2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af774656abc9da3c4c5980cba246bee93bc052c43209b28271259ee431b7fea3"></a>

## http_receiver.use_tls.no_ca — http_receiver.use_tls.no_ca / 3ec358911f71 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-c2515e68697d658421a149636ad34121ab8c5447b571ca1fa29e974052ad74a8)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- http_receiver.use_tls.no_ca

<a id="canonical-5c2bf436a8a202cf82b7c978d7f29587d4cc2448a43ce9ecedc6c1217541c2e1"></a>

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

<a id="canonical-ac9bd65081f108a4d6e0f6e3ec2095e293b4deb9907ad1dc5353d889d82b7062"></a>

## Direct properties — http_receiver.use_tls.no_ca / 3ec358911f71 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b175bc69dfe279f61aebdf2fd88258f9a27cd18dc49829aa33b39e2408df943"></a>

## Next pages — http_receiver.use_tls.no_ca / 3ec358911f71 / 4

- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-ce47dde22abd25c19a84cfcf4cc203acd7ad2a675148eb9430a25cc4168314bc)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2745d9178ffeeabfc18ea105b7d6b625b5f869d75835d804181ec8e35200bb49"></a>

## kafka_receiver — kafka_receiver / 626308e4b80c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- kafka_receiver

<a id="canonical-a31e1cd1ccf0f2715372a58f5a0fc58c6a79d744538de04ea640d384c8395979"></a>

Type: `"single"`. Computed.

Kafka Configuration for Global Log Receiver.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-e277c7ccac32a313215dc59ad000f452233384691407b3e99b387e53312cdb84"></a>

## Direct properties — kafka_receiver / 626308e4b80c / 3

- [batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd): complete subsection reference.

<a id="canonical-7e2cd0638734932dbfa0adfbeddf7c4d5c1911d62517780cabad8dc84eb5aa7d"></a>

<a id="canonical-60586c3609d060720da42652f531fa4de1467cb8caced53e93c8457fa72c32be"></a>

## bootstrap_servers property — kafka_receiver / 626308e4b80c / 4

Type: `["list", "string"]`. Computed.

List of host:port pairs of the Kafka brokers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba): complete subsection reference.

<a id="canonical-86fd8dd45d23424f0efec599305967471a959ee70fee186b837ae6692ae91f1b"></a>

<a id="canonical-7c09ed35e59498b38164459a22f80ac03ca5cf61bbe38011f9436f638d879593"></a>

## kafka_topic property — kafka_receiver / 626308e4b80c / 5

Type: `"string"`. Computed.

The Kafka topic name to write events to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  }
}
```

- [no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-cc9703d87b5cab7e4b8f2bf4ef82c3277817cfd6ce10b5c73ad722a507a46e91): complete subsection reference.

- [use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24): complete subsection reference.

<a id="canonical-6c092e906cf5015bfc6a63cb10d64472df130a10b70de68067b37a94aff42fbd"></a>

## Next pages — kafka_receiver / 626308e4b80c / 6

- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd)
- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba)
- [kafka_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-cc9703d87b5cab7e4b8f2bf4ef82c3277817cfd6ce10b5c73ad722a507a46e91)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9c716a5f78765d821317dcad73bedc9935a9ebf20b7e787116525ce02741a4c"></a>

## kafka_receiver.batch — kafka_receiver.batch / 60f5cd065768 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- kafka_receiver.batch

<a id="canonical-53e9bf94b6c6aac7ccef2340cf936c3620f2979dc650a42e8905dbc602150071"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-ac4f9f196c0a03e633139b8d9c2b4ec7368dc40a7e1e8fcc2eff0a1a6fcbbd54"></a>

## Direct properties — kafka_receiver.batch / 60f5cd065768 / 3

<a id="canonical-86f08a44657eb6e76050db3089764ed44e263dcb4d865ae555c3154ccb1d4dcc"></a>

<a id="canonical-80b38109061498618991f83d005a84bb5dd72ed877a6e0ad0f6724e568a80be3"></a>

## max_bytes property — kafka_receiver.batch / 60f5cd065768 / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-691ce3e3a8e2b029515583c4550e8fcb25eaebfaba1f77d7ab813e2d9d6e4046): complete subsection reference.

<a id="canonical-f3bb912c262ac130008fbb167c6a32075ec2263ef669be881d7b26e4f448ae9d"></a>

<a id="canonical-bce4806e925cbfed245f4acc3bbff1e7790cc283aac760b3fc9e9aef047367c4"></a>

## max_events property — kafka_receiver.batch / 60f5cd065768 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-2cba50c703ea7f01dce3e31854b502d91ff540d34dd9bf97a83d5b8a2b22bf6b): complete subsection reference.

<a id="canonical-b5d48fff30e6b4320b6769c66324bcac0cb4611e29d4fa7fff3b587ecc02dfee"></a>

<a id="canonical-64386c63c37fb90091b4a6fd7b607fcecb0039e2ad0ba938257878f5fa581ddd"></a>

## timeout_seconds property — kafka_receiver.batch / 60f5cd065768 / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-c6c834d83ffa19b0c08be4505a091ce9db369cf7b45dbcc46b6c89bd8c12b882): complete subsection reference.

<a id="canonical-0e87d1ac1425c6bb891d42254856ad4468476085af62119d419ffae75cc291ef"></a>

## Next pages — kafka_receiver.batch / 60f5cd065768 / 7

- [kafka_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-691ce3e3a8e2b029515583c4550e8fcb25eaebfaba1f77d7ab813e2d9d6e4046)
- [kafka_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-2cba50c703ea7f01dce3e31854b502d91ff540d34dd9bf97a83d5b8a2b22bf6b)
- [kafka_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-c6c834d83ffa19b0c08be4505a091ce9db369cf7b45dbcc46b6c89bd8c12b882)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-691ce3e3a8e2b029515583c4550e8fcb25eaebfaba1f77d7ab813e2d9d6e4046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0470577909168abc906ecf80208d25c84193736972997e998905b5cffa214538"></a>

## kafka_receiver.batch.max_bytes_disabled — kafka_receiver.batch.max_bytes_disabled / 707ba382e1fe / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd)
- kafka_receiver.batch.max_bytes_disabled

<a id="canonical-d6c9ec7b49f0b2a215bbff623073a3424c29c7f2cfadec61b35f7a0e2361da69"></a>

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

<a id="canonical-a9a7070c8a0a8eb17f21d545508126a0dbdf8a7def93a49813d6d184266318c8"></a>

## Direct properties — kafka_receiver.batch.max_bytes_disabled / 707ba382e1fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a7dc9b0269d89d61dae470a1a090ab4fa922f6ecf0776e77c47821414cd7286"></a>

## Next pages — kafka_receiver.batch.max_bytes_disabled / 707ba382e1fe / 4

- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-2cba50c703ea7f01dce3e31854b502d91ff540d34dd9bf97a83d5b8a2b22bf6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bfd69976683747ccbee1747c24c0d82acb178797093de12eecfd19dfc92f3fe"></a>

## kafka_receiver.batch.max_events_disabled — kafka_receiver.batch.max_events_disabled / c2d8ee5f798f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd)
- kafka_receiver.batch.max_events_disabled

<a id="canonical-fc83403393425e1344f73dc2a2421b1ce58238060dc9cf3e7113e461e84e0308"></a>

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

<a id="canonical-c32a6f8b7c156924ef319d1fd8c916c6e21f2601141bda57ecdf3811d54e6bc2"></a>

## Direct properties — kafka_receiver.batch.max_events_disabled / c2d8ee5f798f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6ac849accad30340243605c29fcb232504d8adf730e5bcb5cc529f8a08e64d8"></a>

## Next pages — kafka_receiver.batch.max_events_disabled / c2d8ee5f798f / 4

- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c6c834d83ffa19b0c08be4505a091ce9db369cf7b45dbcc46b6c89bd8c12b882"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bbf8fcfb200122d6b04eb4bab697b70fe732629da0682aa2d9f5cca77367049"></a>

## kafka_receiver.batch.timeout_seconds_default — kafka_receiver.batch.timeout_seconds_default / 2b990b4848ed / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd)
- kafka_receiver.batch.timeout_seconds_default

<a id="canonical-de9d06e1d3dad50b6982f4df6af621115325fe21701277f10098f70a0e9325b3"></a>

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

<a id="canonical-4bf02e59bd016a90eaa2b845b16f5e47c66cefd6cc894324092951586501c291"></a>

## Direct properties — kafka_receiver.batch.timeout_seconds_default / 2b990b4848ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-837cfd6096930a79cd849dfa73adb66950f4f34da03dfdf6ea3ba1fe10d3ecb1"></a>

## Next pages — kafka_receiver.batch.timeout_seconds_default / 2b990b4848ed / 4

- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-9157a5b3583a9f094de217dff66d136ad5403d6aaddfb94a8eb0ea6ddeca21dd)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52306e3113234e32808f4fc8a776a2c3a6659124fe174180b027931d887628f8"></a>

## kafka_receiver.compression — kafka_receiver.compression / 300b7114a6df / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- kafka_receiver.compression

<a id="canonical-dadf47410aa391ddf47afa8ea17ab1746f66e8666ba9358bf26aa3bb7df367e8"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-e8ae7b070b200fce79ff9be824bdc12f28adbeec390e8d0554c235f149a25a93"></a>

## Direct properties — kafka_receiver.compression / 300b7114a6df / 3

- [compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-c521ae21b95049a1e552b8465b7c849bc4c03563c82893b5a6321b63396f28b2): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-f09f49473ca6ccf5085ecf1248e73f835b06c81158c45d31e77ff63c60ce5f24): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-3ff85e1efc877a8d8ac5a8742a4690f5de39fde7e41656a951aa4f483de083ee): complete subsection reference.

<a id="canonical-d720c4ac2893364374461b1092dbd3f6798430ec15fb46474d9d655da6e60dc2"></a>

## Next pages — kafka_receiver.compression / 300b7114a6df / 4

- [kafka_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-c521ae21b95049a1e552b8465b7c849bc4c03563c82893b5a6321b63396f28b2)
- [kafka_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-f09f49473ca6ccf5085ecf1248e73f835b06c81158c45d31e77ff63c60ce5f24)
- [kafka_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-3ff85e1efc877a8d8ac5a8742a4690f5de39fde7e41656a951aa4f483de083ee)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c521ae21b95049a1e552b8465b7c849bc4c03563c82893b5a6321b63396f28b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80603b3b6d5fd1aaeb81ae7af95cddd12dbd1ede7869997977bda87b31217a54"></a>

## kafka_receiver.compression.compression_default — kafka_receiver.compression.compression_default / 1a812b391451 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba)
- kafka_receiver.compression.compression_default

<a id="canonical-c8436c1f40229d3a3f12a43978b664f211275957632585b6ba8b39a227103813"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-4b80f7dc5c964707b4a9a574bad492148bb922ec94a0972c6b03cfe27b55e593"></a>

## Direct properties — kafka_receiver.compression.compression_default / 1a812b391451 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffce50a6018449b8b9955b1b906059614835e5ba9fa12792985f76fe898b2e36"></a>

## Next pages — kafka_receiver.compression.compression_default / 1a812b391451 / 4

- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f09f49473ca6ccf5085ecf1248e73f835b06c81158c45d31e77ff63c60ce5f24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-891827dc8b280c45e93f31eb43c6e88b15d5467dc6100fa76e64390a4ee78cbb"></a>

## kafka_receiver.compression.compression_gzip — kafka_receiver.compression.compression_gzip / fc9b433383ee / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba)
- kafka_receiver.compression.compression_gzip

<a id="canonical-855b25b4f3642ff8a85dd5eec79358bba4c81b03002a9642d891b501a703f605"></a>

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

<a id="canonical-1afede6a038be2d35c888f9dbf6e9ff6f296050477ce72e03cd57ab1681c383d"></a>

## Direct properties — kafka_receiver.compression.compression_gzip / fc9b433383ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a2bb3e896f3d54ffa31441a8e140dd4a0f37f4e4119ecba4493835afc129893"></a>

## Next pages — kafka_receiver.compression.compression_gzip / fc9b433383ee / 4

- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-3ff85e1efc877a8d8ac5a8742a4690f5de39fde7e41656a951aa4f483de083ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed56167a9c59b39eb6dc6240b05e8119113d821e1b60fe7b615f19ed80fd92e7"></a>

## kafka_receiver.compression.compression_none — kafka_receiver.compression.compression_none / 064901a0c56b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba)
- kafka_receiver.compression.compression_none

<a id="canonical-e1eac692a08e188a817fb4ff376811de7723fa6d4e6c6d0d254dc311bed66203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-282e62f5cad5b875d0bdcdcdec52188b4cb2ded1b653cafdaf3dd026655f2ff0"></a>

## Direct properties — kafka_receiver.compression.compression_none / 064901a0c56b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f36b02c17f80e8b86f54d368c32880616ccdcfbd96b59cb3978227a018faf98"></a>

## Next pages — kafka_receiver.compression.compression_none / 064901a0c56b / 4

- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-ba97208b29e3c438f52532d617c628b20b82a85b66076dfe694258789f4c71ba)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-cc9703d87b5cab7e4b8f2bf4ef82c3277817cfd6ce10b5c73ad722a507a46e91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77a1d01402f9127fac143b37bbece13b74e9c6ba80d71adf419ef245cad99cbe"></a>

## kafka_receiver.no_tls — kafka_receiver.no_tls / 58fd356e60b4 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- kafka_receiver.no_tls

<a id="canonical-bd04ced755f2bd79e01964c4ff110b46efc7bef0e198a24c4c13a851266cea4a"></a>

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

<a id="canonical-81502cc5fe897a69b874a96e2233d36250c5a7eadf5577a44328669495675c19"></a>

## Direct properties — kafka_receiver.no_tls / 58fd356e60b4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc7b2707e048802f2318f08031d55b011e9f076a7258d3af5ae74be0fd6ec57d"></a>

## Next pages — kafka_receiver.no_tls / 58fd356e60b4 / 4

- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30defd0f851b8d51087979fa93595d1ab504ef248778865d23ee7891ddb92697"></a>

## kafka_receiver.use_tls — kafka_receiver.use_tls / 25d79f971ee1 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- kafka_receiver.use_tls

<a id="canonical-fc7a0410084a7323afbf3d3432f45fa98857d4cd46fea4673d8314262ebe8d7d"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

<a id="canonical-c9e9a799ebeac9cadfd45d084d285392abec472c39214fa445e639af68d4e053"></a>

## Direct properties — kafka_receiver.use_tls / 25d79f971ee1 / 3

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-97296f87ca7d9afef58d86f91639112b9f04f3d133c9b926ddd6d6d924e1b2b4): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-51a4579aa50463d6f75767919fe0fc15df1c0922cf598474c50111b77ed79a9d): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-a158032a2bb89062eb719838dd839afb7f2fb374f681ae4d29bc9586e95ddc07): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-73922908c9c2f5c693d5ab901e19970258315f6068af0cc8d1199b4390effbcb): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-38694bb626dfd35cad1ac518058fcb5a88b67c7ec9833e820499216787d2798d): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-61b57298607e1d82c2c7b0d98f9d896c1da1c3109d3a3ffaf6979a392b580f84): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-6ed65979fa0f3c51b0652925afef73cf55ea7691a607613655da0d706181520b): complete subsection reference.

<a id="canonical-2d8f708a6168f399bb8e044c07091bc1aff3895282c30c4a2cc6d7a473c21274"></a>

<a id="canonical-175aafc29c8a2cc92fa7a4e0dc757fd9e6e4ae2176f1874a2ba0e16b926454c9"></a>

## trusted_ca_url property — kafka_receiver.use_tls / 25d79f971ee1 / 4

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-dfb98852ac87e7cc9a206c3b72950ec81f90cf744b057f2267fa6d7cabe2ad8e"></a>

## Next pages — kafka_receiver.use_tls / 25d79f971ee1 / 5

- [kafka_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-97296f87ca7d9afef58d86f91639112b9f04f3d133c9b926ddd6d6d924e1b2b4)
- [kafka_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-51a4579aa50463d6f75767919fe0fc15df1c0922cf598474c50111b77ed79a9d)
- [kafka_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-a158032a2bb89062eb719838dd839afb7f2fb374f681ae4d29bc9586e95ddc07)
- [kafka_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-73922908c9c2f5c693d5ab901e19970258315f6068af0cc8d1199b4390effbcb)
- [kafka_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-38694bb626dfd35cad1ac518058fcb5a88b67c7ec9833e820499216787d2798d)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-61b57298607e1d82c2c7b0d98f9d896c1da1c3109d3a3ffaf6979a392b580f84)
- [kafka_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-6ed65979fa0f3c51b0652925afef73cf55ea7691a607613655da0d706181520b)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-97296f87ca7d9afef58d86f91639112b9f04f3d133c9b926ddd6d6d924e1b2b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbea724451f71c672a5f929f01b97b4b229e39b2949edef45d14a5c9b5385207"></a>

## kafka_receiver.use_tls.disable_verify_certificate — kafka_receiver.use_tls.disable_verify_certificate / 951864876ec5 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- kafka_receiver.use_tls.disable_verify_certificate

<a id="canonical-6b869ded0d6fd18c4ce1882aa8db864cbfce7c34e19782fcadb8f106a2c898b0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-82b465e96f6691c6af6a1315d287636bb3c9000dc87994ead51cbb13d14d5fb2"></a>

## Direct properties — kafka_receiver.use_tls.disable_verify_certificate / 951864876ec5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d256bec61a9c10998daca2f78ccd0c28e7bc0a8e3dd3bdad486d1684c5607602"></a>

## Next pages — kafka_receiver.use_tls.disable_verify_certificate / 951864876ec5 / 4

- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-51a4579aa50463d6f75767919fe0fc15df1c0922cf598474c50111b77ed79a9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3d92554d6a40b2571c4aea3e681474efa9ac02eeeefc400502a8c2c97db5e3c"></a>

## kafka_receiver.use_tls.disable_verify_hostname — kafka_receiver.use_tls.disable_verify_hostname / 0ac77bd0ca9c / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- kafka_receiver.use_tls.disable_verify_hostname

<a id="canonical-4909e83b6a5c710921b55c398053237589737feea8831ac90ad0396db844a40f"></a>

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

<a id="canonical-21392312bef135e98ce3591fbe5144dc9276aef4b891d6f9c1553fb42cb33381"></a>

## Direct properties — kafka_receiver.use_tls.disable_verify_hostname / 0ac77bd0ca9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78237fc11d17e63d61a355bcda725ec094148736215a3909e4b297b64de9863e"></a>

## Next pages — kafka_receiver.use_tls.disable_verify_hostname / 0ac77bd0ca9c / 4

- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a158032a2bb89062eb719838dd839afb7f2fb374f681ae4d29bc9586e95ddc07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-206c9659b7003ed14f2b3d3540ac0b7bdacc5e8674513e223afba7df9901374a"></a>

## kafka_receiver.use_tls.enable_verify_certificate — kafka_receiver.use_tls.enable_verify_certificate / 8ded4259df2f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- kafka_receiver.use_tls.enable_verify_certificate

<a id="canonical-539556431e452db7e8d30dbbc66125b08ec40c9d2c2a7a67ff731dde9395cab2"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-dcafc97f56db47294a63b4be985ef155dd3292c26ce366bb5050ced0038b773a"></a>

## Direct properties — kafka_receiver.use_tls.enable_verify_certificate / 8ded4259df2f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b119de8c61d0cf99ce839fa568df654a4c54ebc8f7b639b4e9243b44db8db56"></a>

## Next pages — kafka_receiver.use_tls.enable_verify_certificate / 8ded4259df2f / 4

- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-73922908c9c2f5c693d5ab901e19970258315f6068af0cc8d1199b4390effbcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59543cca57912214cd4619a2fbf7ddaa8525a6e09a2a7628f33a32bd0b041362"></a>

## kafka_receiver.use_tls.enable_verify_hostname — kafka_receiver.use_tls.enable_verify_hostname / 49469913adc8 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- kafka_receiver.use_tls.enable_verify_hostname

<a id="canonical-f21707b40187ac8e273a44b8e64aab24dabe9bb8cffabc08b3f7afbfa14e32a4"></a>

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

<a id="canonical-7f48f90dc0a3d6e83e936d4b9bb2a3b184f85550574e07c5152861f92a010b3a"></a>

## Direct properties — kafka_receiver.use_tls.enable_verify_hostname / 49469913adc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-99f9f19ba96e1ef6a58bb3428ead2114f471a7ea6fedfd9309c80e8e884be799"></a>

## Next pages — kafka_receiver.use_tls.enable_verify_hostname / 49469913adc8 / 4

- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-38694bb626dfd35cad1ac518058fcb5a88b67c7ec9833e820499216787d2798d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-091d8b652870db4ce48fa0213edda6eb72f6382712564d9970321358adda0db8"></a>

## kafka_receiver.use_tls.mtls_disabled — kafka_receiver.use_tls.mtls_disabled / ae6780e7bda7 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- kafka_receiver.use_tls.mtls_disabled

<a id="canonical-86403a836f2b7c58623b124f9df0fc74b080094ce22d379f838616ee66c32717"></a>

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

<a id="canonical-10eb5bc7742cd4bfc5dd6186f63a7e32068f9e1a2eddc56679e2f7299ad5fb5a"></a>

## Direct properties — kafka_receiver.use_tls.mtls_disabled / ae6780e7bda7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-400f80e05dc9f7518433331ba522870bea8204bcded688c4765fa3c189148088"></a>

## Next pages — kafka_receiver.use_tls.mtls_disabled / ae6780e7bda7 / 4

- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-61b57298607e1d82c2c7b0d98f9d896c1da1c3109d3a3ffaf6979a392b580f84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0d028e4e4b4e9060d747c854a1aba65c7119fede269b519d716fb615446934f"></a>

## kafka_receiver.use_tls.mtls_enable — kafka_receiver.use_tls.mtls_enable / 9be5edd64ef2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- kafka_receiver.use_tls.mtls_enable

<a id="canonical-4f82aee5b8986387dd0a106587743aae6329f8c08936de64966b0e84a2099b22"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

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

<a id="canonical-35315c6b208a6e49a5a74e29e7ed7c118776c4d3007bcba091360995bed620f7"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable / 9be5edd64ef2 / 3

<a id="canonical-1a6b774764971831650f041835945bb5a33d8ee8238530189e5db1b53dbd0092"></a>

<a id="canonical-25e952d8978671e32c42f155d5929f8f158faaf55d02e5a605544456d617521d"></a>

## certificate property — kafka_receiver.use_tls.mtls_enable / 9be5edd64ef2 / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-aa695d207478a160fbf9015d675930894fce429844c84de69ba981456ccda865): complete subsection reference.

<a id="canonical-f748ab148792ea97a4f9a81605b255895ba89fbd3afda9aa2a163710ab5084b7"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable / 9be5edd64ef2 / 5

- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-aa695d207478a160fbf9015d675930894fce429844c84de69ba981456ccda865)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-aa695d207478a160fbf9015d675930894fce429844c84de69ba981456ccda865"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b95cb234e220208ca3e8f3dea2997ba521643b5817ab4f2afa91c8119297b2af"></a>

## kafka_receiver.use_tls.mtls_enable.key_url — kafka_receiver.use_tls.mtls_enable.key_url / 1100409c89d2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-61b57298607e1d82c2c7b0d98f9d896c1da1c3109d3a3ffaf6979a392b580f84)
- kafka_receiver.use_tls.mtls_enable.key_url

<a id="canonical-f56b685179c49e34e8efc65c61cf6adbf0dbb9d1133fdb8a6f079dfcad45856f"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-dfe3c88e7250b06bb123bf7d3af36209d2852c9a8648a7258b09c6719b363b93"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable.key_url / 1100409c89d2 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-05a2e4abd81b17742575a80d62015f57f15b6e4f19281c437b6d3357ba9cf064): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-726a7490d8789c48923647bfe0301bb57f08f369bc1a8aee83aaefb1ad336857): complete subsection reference.

<a id="canonical-1f08beb52824e35f33abf1577ee7593b8e6ec4848938dc406aebe3a9900dc706"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable.key_url / 1100409c89d2 / 4

- [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-05a2e4abd81b17742575a80d62015f57f15b6e4f19281c437b6d3357ba9cf064)
- [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-726a7490d8789c48923647bfe0301bb57f08f369bc1a8aee83aaefb1ad336857)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-61b57298607e1d82c2c7b0d98f9d896c1da1c3109d3a3ffaf6979a392b580f84)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-05a2e4abd81b17742575a80d62015f57f15b6e4f19281c437b6d3357ba9cf064"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c6fc639c4b10bc992e7d42bc1b08d3c46ca5f7c5d95354ee30659f698edef71"></a>

## kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 38f3f118c3fc / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-61b57298607e1d82c2c7b0d98f9d896c1da1c3109d3a3ffaf6979a392b580f84)
- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-aa695d207478a160fbf9015d675930894fce429844c84de69ba981456ccda865)
- kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0516261b0949d016e4fa60764d64a1d836f82a4a0b729d3d02c23152ca1eb885"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-f4fc72108a0820ea0f0714434277f33b74e7aeeff4ed083b538425c473acc931"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 38f3f118c3fc / 3

<a id="canonical-6a598e77b7343bdbb07574676628c5ab0ba54ddfd2c755a9695045f87ec9d465"></a>

<a id="canonical-fab4fa856f12bcc0f9c63387ae6459c5d517c054140c4b9727abedc2814f9f33"></a>

## decryption_provider property — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 38f3f118c3fc / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-a40f121f3ff7e7cccdf70ef25fb2428d593a2048486fa2117278d4ffb0e657dc"></a>

<a id="canonical-d3e98a3e35729aea9ee771cc233ba825d0b2f31908541884d7939b6494d02831"></a>

## location property — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 38f3f118c3fc / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-b3a4a19f341cdeeab62fb551e37776ed7ff80c94ae19c30f686dbeec2e64a344"></a>

<a id="canonical-d6f0b7edaa1bdcc8fe9b94176883b5f25df7f3b2a739a804cf5d344191baa4e0"></a>

## store_provider property — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 38f3f118c3fc / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-e641d34b67104d1db3c4119947532579a902d1bfd9d72ed54166e0378ac51f95"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info / 38f3f118c3fc / 7

- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-aa695d207478a160fbf9015d675930894fce429844c84de69ba981456ccda865)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-726a7490d8789c48923647bfe0301bb57f08f369bc1a8aee83aaefb1ad336857"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51ea959ad73816273dc9286ce341a18311f503bfce80739769cc12b4a7348639"></a>

## kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / a774fa65af0e / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-61b57298607e1d82c2c7b0d98f9d896c1da1c3109d3a3ffaf6979a392b580f84)
- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-aa695d207478a160fbf9015d675930894fce429844c84de69ba981456ccda865)
- kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-a114592615df6133f0aaf34c9cf40ef0cf88280f732b7b3f58f032af81d53bf0"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-e69ae31717e897e008ba647e9108125f489e5be44593294931430bc11bf303f6"></a>

## Direct properties — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / a774fa65af0e / 3

<a id="canonical-99584070aa26d4031e82ef1f5565a405f0cf4fa6a0f990c676608d2cbe19f359"></a>

<a id="canonical-01d7c1bd1ab06ec9bb921f38e343676f4d5924e043cabfd768108d0e5d80f793"></a>

## provider_ref property — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / a774fa65af0e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8c76f0a6989b2634fb5c8a0d92bf72f158d6f45dd36f112d3be7b48adb977d63"></a>

<a id="canonical-1f6bbb7d4137750f9cb1ec4c6169d7757cf336132a21cc906685cdbd23493585"></a>

## url property — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / a774fa65af0e / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-014a23df80ba8b0557adc35a4e6377a88f0d238b5bbc8c6505c6e3ccb5cde9bd"></a>

## Next pages — kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info / a774fa65af0e / 6

- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-aa695d207478a160fbf9015d675930894fce429844c84de69ba981456ccda865)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6ed65979fa0f3c51b0652925afef73cf55ea7691a607613655da0d706181520b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-780fca3e2cc08522aa3da56e94259b82ff2f768a17ee68faabc7d09229ea0c1e"></a>

## kafka_receiver.use_tls.no_ca — kafka_receiver.use_tls.no_ca / 52a019591df1 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b2ba99714bf357739d2e73a2af07b4be508bfcec644db2fba3f3e59a26fa81bd)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- kafka_receiver.use_tls.no_ca

<a id="canonical-a3c91c6c718fd42b4e06946653535eafcef04e43cc5cccaaa372a0bfd915e652"></a>

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

<a id="canonical-b35a3ecccb8c7a0a5f06f86ca4f885ecfa33a169cf19828085dfc18d5ece6599"></a>

## Direct properties — kafka_receiver.use_tls.no_ca / 52a019591df1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d8919c6a69a5551f5c02f7ca92d3eaf3d349ff19f98625096157255764c8b85e"></a>

## Next pages — kafka_receiver.use_tls.no_ca / 52a019591df1 / 4

- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a28eb56bde911fa8f8ab4623ed60497d2de2fbad488e6c288ab10fdcc56f8a24)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79a43a40c824d0d95f11b7afdb69e9b633946dce4b9fb2436a90ae64478913be"></a>

## new_relic_receiver — new_relic_receiver / 875db0c75830 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- new_relic_receiver

<a id="canonical-de4a336b302a3b488fb87be5ddc2114a24a5cc1ba754c0ed424c2c8c71365141"></a>

Type: `"single"`. Computed.

Configuration parameter for new relic receiver.

Upstream description:

Configuration for NewRelic endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"eu\",\"us\"]"
}
```

<a id="canonical-0693e6f55aa724ef33ab7946ae1a244d4e51cedf6813c48c4c2f6f7cdaa8b6a1"></a>

## Direct properties — new_relic_receiver / 875db0c75830 / 3

- [api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-a5d111e2ab04bd3bcdd2a3b8efbd88ad11e9af145eb8f0f72cf3a32627456133): complete subsection reference.

- [eu](data-sources--global_log_receiver--reference--group-003.md#canonical-f0126582e7b7d779bd52ca048638dfff39b2c5cc9543b207d7986bec144c42d7): complete subsection reference.

- [us](data-sources--global_log_receiver--reference--group-003.md#canonical-c6ce4697b79b9d0bde57ec9e05ee4a32d37851f1c9439a78a387b45282378086): complete subsection reference.

<a id="canonical-d7b9e6dc175f9c2f8247705ccb46cd8a0a7fa52fc6aafc5c67300eb99e6fcade"></a>

## Next pages — new_relic_receiver / 875db0c75830 / 4

- [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-a5d111e2ab04bd3bcdd2a3b8efbd88ad11e9af145eb8f0f72cf3a32627456133)
- [new_relic_receiver.eu](data-sources--global_log_receiver--reference--group-003.md#canonical-f0126582e7b7d779bd52ca048638dfff39b2c5cc9543b207d7986bec144c42d7)
- [new_relic_receiver.us](data-sources--global_log_receiver--reference--group-003.md#canonical-c6ce4697b79b9d0bde57ec9e05ee4a32d37851f1c9439a78a387b45282378086)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a5d111e2ab04bd3bcdd2a3b8efbd88ad11e9af145eb8f0f72cf3a32627456133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-510d098b1ef991002d8c3c5835031ea5e6696616b65dae5db9f967142fbe9c96"></a>

## new_relic_receiver.api_key — new_relic_receiver.api_key / 40d966a85bcf / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- new_relic_receiver.api_key

<a id="canonical-70959860813e842268b5c73806c213433e44710f27ee47f4b5ae4c8cad1dc771"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-b478c8202d1e68c03adf1ea335b77114ec68728a31004e7b86f18fa9f475a5fc"></a>

## Direct properties — new_relic_receiver.api_key / 40d966a85bcf / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-1b05515347110c5166efb415fd2d1e49cad46c5e39cf44bf01266a6484523191): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-78a657c94fc131985633c3a14d87e6d17b811c07aa1a4648a0713d638ebba33d): complete subsection reference.

<a id="canonical-6d91c6461edd2e248b4a750d65c09dbe3bdc5f9b761ed0e7a365ba00139b8bc5"></a>

## Next pages — new_relic_receiver.api_key / 40d966a85bcf / 4

- [new_relic_receiver.api_key.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-1b05515347110c5166efb415fd2d1e49cad46c5e39cf44bf01266a6484523191)
- [new_relic_receiver.api_key.clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-78a657c94fc131985633c3a14d87e6d17b811c07aa1a4648a0713d638ebba33d)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-1b05515347110c5166efb415fd2d1e49cad46c5e39cf44bf01266a6484523191"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc04ca7691835c594e0c97574f3f2dc2f652449809a5f38e31087a32a165b688"></a>

## new_relic_receiver.api_key.blindfold_secret_info — new_relic_receiver.api_key.blindfold_secret_info / f7a119019315 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-a5d111e2ab04bd3bcdd2a3b8efbd88ad11e9af145eb8f0f72cf3a32627456133)
- new_relic_receiver.api_key.blindfold_secret_info

<a id="canonical-6d4f84f71264dacc88d3fc0abbffe91102fb1d2c1b984707656c3d035ed2f383"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-9f3ba553a98631acc5f77395cfaf937cd69f801c518de74ff3f2732054b73d27"></a>

## Direct properties — new_relic_receiver.api_key.blindfold_secret_info / f7a119019315 / 3

<a id="canonical-de82c3b32fc5fa2bce3ab23d8b5c5cbb06e83d47aab0a6fa519d8b8768e4434a"></a>

<a id="canonical-63784d4e0ef56e6b6560d6665f68a0cc693f805290ee28a34451f4dcf6df98c1"></a>

## decryption_provider property — new_relic_receiver.api_key.blindfold_secret_info / f7a119019315 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0b233eac3d86a24b5c76c0cfc3c38e3e721bd97e5104769a73672286a4fb60c9"></a>

<a id="canonical-3dbe8587fcafabbadcbda9b2dafff063574f8f48dcb665fd94322c3cf13143f4"></a>

## location property — new_relic_receiver.api_key.blindfold_secret_info / f7a119019315 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-b3fa2c78844c49dd7524cd3a32f324c532ddb711162c76f4d469bf09e59074a4"></a>

<a id="canonical-378cd523a34ddb748128482418b23b978e3e0e52d156f371702c5731dfa536ab"></a>

## store_provider property — new_relic_receiver.api_key.blindfold_secret_info / f7a119019315 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-e67b2cc701edcd1bfcc4518c3df99bb15fc1ebbd9417be278ab4db03c0bb6f3b"></a>

## Next pages — new_relic_receiver.api_key.blindfold_secret_info / f7a119019315 / 7

- [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-a5d111e2ab04bd3bcdd2a3b8efbd88ad11e9af145eb8f0f72cf3a32627456133)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-78a657c94fc131985633c3a14d87e6d17b811c07aa1a4648a0713d638ebba33d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bb42069cd5e2f637f42c4d5826e318cf4b78772af1a9703d01b628258b3ed5e"></a>

## new_relic_receiver.api_key.clear_secret_info — new_relic_receiver.api_key.clear_secret_info / bce3b91958b3 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-a5d111e2ab04bd3bcdd2a3b8efbd88ad11e9af145eb8f0f72cf3a32627456133)
- new_relic_receiver.api_key.clear_secret_info

<a id="canonical-bbef8af3f744997fb19a474aa17488cfd6d35fdee771b8b394c45f6b76326141"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-f896b00dab20917d6d0c4463c8105cf229c588771706994600ad291e5df5a24c"></a>

## Direct properties — new_relic_receiver.api_key.clear_secret_info / bce3b91958b3 / 3

<a id="canonical-0993bb241a1066539e70aea7adc033ef2150b5f5db3c07fb8b1daf3dae4e0a09"></a>

<a id="canonical-354bbb40fec2802866c22392dc7f355eb59314c6db7f8e64bf8b9663114c1b6d"></a>

## provider_ref property — new_relic_receiver.api_key.clear_secret_info / bce3b91958b3 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-656a13b8277ed158fe98c12acc1d34f32d3e349aed2df172f5b62e7e3fb3b9a6"></a>

<a id="canonical-aa3756e026457642089d5b433bdc91157c4f9cdeb672f00e7855f4fc0b3ac16f"></a>

## url property — new_relic_receiver.api_key.clear_secret_info / bce3b91958b3 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-d30f51dc91ebef7edd663b4d4d2cab3b110109d4976a00dc9175557bca478878"></a>

## Next pages — new_relic_receiver.api_key.clear_secret_info / bce3b91958b3 / 6

- [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-a5d111e2ab04bd3bcdd2a3b8efbd88ad11e9af145eb8f0f72cf3a32627456133)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f0126582e7b7d779bd52ca048638dfff39b2c5cc9543b207d7986bec144c42d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a730ba8e98d0fe022e36479f15fc8ab6a40cfe67939490edd47368a2e7f3cba1"></a>

## new_relic_receiver.eu — new_relic_receiver.eu / 3b981474270b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- new_relic_receiver.eu

<a id="canonical-a97238c0fe6143038aef5dcb1f5cb04b653ef091e11b6bfc311a9bd564666e2a"></a>

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

<a id="canonical-51fd64615a5836cdc40d47103ada8667a886b21d6b52bdf5c14d3141e37ab096"></a>

## Direct properties — new_relic_receiver.eu / 3b981474270b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6cae528abfc93a0f1e5b31c97400729eeb31a003f0a242ea48d2bbd94b33abf3"></a>

## Next pages — new_relic_receiver.eu / 3b981474270b / 4

- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c6ce4697b79b9d0bde57ec9e05ee4a32d37851f1c9439a78a387b45282378086"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44fcaa7056dd15f1fe88eef04bc9cc48ede33a656dc7023fa0401cdf8e4acd42"></a>

## new_relic_receiver.us — new_relic_receiver.us / 9c97b454af6a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- new_relic_receiver.us

<a id="canonical-f267c9272a349908094306debf89123593f1f52c89f843ed9ef97ba7a0209904"></a>

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

<a id="canonical-9de0a10f61dc4b8624d9143b6f4a6f94e295f38c05058187c284cbd4b5c54bae"></a>

## Direct properties — new_relic_receiver.us / 9c97b454af6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a621aef2ae136f6e082259496938401c1ea494e6439d043ebd2455911c4c3251"></a>

## Next pages — new_relic_receiver.us / 9c97b454af6a / 4

- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-6a8ca15a6755e8c761d98fa0efc7d887c67f7c2cf83ede5ae159146cee36621c)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-eac6c2cc015966b8d0a1b560e162a5757c3992ef61ab6e3555f7a8641fdb1d88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22822cdd05f63ffc98d1058d008027c5d6dc2cae3c9f3b4627d4276e15b084f4"></a>

## ns_all — ns_all / 45f9655ac63f / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- ns_all

<a id="canonical-4c61613dac1b7423de7e703f6ebfc191efeb081005fa60f1948e50325101b0c7"></a>

Type: `["object", {}]`. Computed.

\[OneOf: ns\_all, ns\_current, ns\_list\] Enable this option

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

- [ns_all](data-sources--global_log_receiver--reference--group-003.md#canonical-4c61613dac1b7423de7e703f6ebfc191efeb081005fa60f1948e50325101b0c7)
- [ns_current](data-sources--global_log_receiver--reference--group-003.md#canonical-a0e14c04d927e045973f735ce6e43e5cab4df82b3725fd35a09f0773ca47fc02)
- [ns_list](data-sources--global_log_receiver--reference--group-003.md#canonical-c5020c3ff01ccbcb7f51e4dd0018c982755920a0d4dfa0b182cbc771682ad584)

Select alternatives according to the provider validators above.

<a id="canonical-440b6c2113873b0bc96a702468307be087fbcc79682e8565237203a83bab2361"></a>

## Direct properties — ns_all / 45f9655ac63f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51f8d534be57d7d7d58cf3b49f339bd06175e28bd5000b1e0586e9eed6d71db1"></a>

## Next pages — ns_all / 45f9655ac63f / 4

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e7d88a8682a21bc20dc8d78ab563c10f0d3a2ee20a4b1399394c1bda0161b874"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07155885b471a83878f672e3771e779f14c1463ccc080a80b630e90cc1ba42ef"></a>

## ns_current — ns_current / 019f476366a2 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- ns_current

<a id="canonical-a0e14c04d927e045973f735ce6e43e5cab4df82b3725fd35a09f0773ca47fc02"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-69f168b6297846812390cd1d37ad1112c93d36f06059dddcf1c22f85c8c457e2"></a>

## Direct properties — ns_current / 019f476366a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1067280cadbdd4e884eb3441a32a82ed3bda7c6dc512740df89429a1bef03e6f"></a>

## Next pages — ns_current / 019f476366a2 / 4

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-f360f7b122a79b7d1c6f6f9d22020aaec9eb1157401d37099f911acc32f3529a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03004c8e34c0c5ef6867ee7d7ad592f33ded8b29fbe5f96fb3e241f9541a5f12"></a>

## ns_list — ns_list / f702ec86e9fa / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- ns_list

<a id="canonical-c5020c3ff01ccbcb7f51e4dd0018c982755920a0d4dfa0b182cbc771682ad584"></a>

Type: `"single"`. Computed.

Namespace List. Namespace List.

Upstream description:

Namespace List.

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

<a id="canonical-40991243b12cf203ef3c7b6c624da85f4dd7bca9ef53fbeeca4bd2a31588d19d"></a>

## Direct properties — ns_list / f702ec86e9fa / 3

<a id="canonical-b7bc4b2fd247f874259644c28d852e75bb063e2e306a68f8aaa03ebbdc7b2546"></a>

<a id="canonical-7f8aa03673c6bea2675608276bb73ad796d0be42902ea2d45910d18249733869"></a>

## namespaces property — ns_list / f702ec86e9fa / 4

Type: `["list", "string"]`. Computed.

Namespaces. List of namespaces to stream logs for.

Upstream description:

List of namespaces to stream logs for.

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

<a id="canonical-a9365f34a460819d42bbdc18817b53f89d923971d6db73577bbbc808e2075a74"></a>

## Next pages — ns_list / f702ec86e9fa / 5

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7376ae09ec115fa389bfe3ed219d9e165c6a93ae175b8902a0f63b97a676774f"></a>

## qradar_receiver — qradar_receiver / 97b36f995835 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- qradar_receiver

<a id="canonical-63e1be9603a2f1112924bcec75ff8b9f9b9b2825fccf9f55c0d1c4a43b8caf74"></a>

Type: `"single"`. Computed.

Configuration parameter for qradar receiver.

Upstream description:

Configuration for IBM QRadar endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-43aeeeef8e8941f3e73ece3c0ba1e1dc787be97deaa327cccff4cacb886c7d02"></a>

## Direct properties — qradar_receiver / 97b36f995835 / 3

- [batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4): complete subsection reference.

- [no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-55ec99baa6f49fbf244a8fbfdba9501b9d3552dc096a2d9bf1d6257c2613a3a6): complete subsection reference.

<a id="canonical-7647ab2b5fccea37a28b2aee902ae6b9a441378cd8f69e288091130b0b71a504"></a>

<a id="canonical-e6d5207fbfddaea977d0e01353fcc870df3354107ea013fb682cedaf0facc1bd"></a>

## uri property — qradar_receiver / 97b36f995835 / 4

Type: `"string"`. Computed.

Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5): complete subsection reference.

<a id="canonical-f191e5dcae2ac5db56391273de981a05f243a0ba19d3127cf5f2b2c1ea664c8f"></a>

## Next pages — qradar_receiver / 97b36f995835 / 5

- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61)
- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4)
- [qradar_receiver.no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-55ec99baa6f49fbf244a8fbfdba9501b9d3552dc096a2d9bf1d6257c2613a3a6)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-362d10d5d0eb99cc2d27325df938568bb84fbfa93aab677eb79e290b0cb3c4b9"></a>

## qradar_receiver.batch — qradar_receiver.batch / 92a52dfeb2ec / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- qradar_receiver.batch

<a id="canonical-18d61cb5d591896c9ae80eb6d9e52e9c210d02fb4dd296d4570162f356e9f8c1"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-e9962825ce7f8b9fbe657e0b90fd9d154811afa7ee3c14148c20435f2ce0d409"></a>

## Direct properties — qradar_receiver.batch / 92a52dfeb2ec / 3

<a id="canonical-f8156477011b20209d669f357f9e64dc712c4a11019b962ac887a02c270d7b03"></a>

<a id="canonical-68287accc41ec0ef9040b4efa5f279b2b8e7487a43aee76edf5a289c6fa828a1"></a>

## max_bytes property — qradar_receiver.batch / 92a52dfeb2ec / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-a07b1b643e8b0b11593399053c70baf7100576425bba064ba6549c8ddf5a410e): complete subsection reference.

<a id="canonical-17a67128efe8c65264683d106db6986d0392430cb7d51e85e996c6225720b1da"></a>

<a id="canonical-5f2adafc021c8c9fbf1ee8e45a80ab0bfd470d2083cfac19ce1aa63d245c317e"></a>

## max_events property — qradar_receiver.batch / 92a52dfeb2ec / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-494d73db92f553ebadb34dbc35913c4b785e42cd0cac3b96fbe67393f6eb0524): complete subsection reference.

<a id="canonical-5106ec2de2aeec40ba6c0f213fcb939f44e9364a89361faeba90177890e78558"></a>

<a id="canonical-07562465cf481574f46a4c2f61a69eedb0007b52980d0b019ffd23583485364f"></a>

## timeout_seconds property — qradar_receiver.batch / 92a52dfeb2ec / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-45bc78c0697d2d94c38d472c9ca67834d55f8162e1810df42a877271b6c165ed): complete subsection reference.

<a id="canonical-a768d0cf3e00bb96716b685100a454d295896648dbdb47f4ad3c13b0600baa09"></a>

## Next pages — qradar_receiver.batch / 92a52dfeb2ec / 7

- [qradar_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-a07b1b643e8b0b11593399053c70baf7100576425bba064ba6549c8ddf5a410e)
- [qradar_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-494d73db92f553ebadb34dbc35913c4b785e42cd0cac3b96fbe67393f6eb0524)
- [qradar_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-45bc78c0697d2d94c38d472c9ca67834d55f8162e1810df42a877271b6c165ed)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a07b1b643e8b0b11593399053c70baf7100576425bba064ba6549c8ddf5a410e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-834d873349988008768fbdfe5aae38943fa933a18ec46628dd1b943fb97cf522"></a>

## qradar_receiver.batch.max_bytes_disabled — qradar_receiver.batch.max_bytes_disabled / a5c211f0cc79 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61)
- qradar_receiver.batch.max_bytes_disabled

<a id="canonical-e2eb3fc8f3b0ff4b61accfd0c69d8638a00216964eadfbfd1bf7753604e34179"></a>

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

<a id="canonical-792b20e71d566d0a5b0f4397998b80a7d9cb2af0232245fad712277ddcbdf302"></a>

## Direct properties — qradar_receiver.batch.max_bytes_disabled / a5c211f0cc79 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-358e6e0476f304ed85ddac73009985eede3e38b68100d1de410ecb8dc80822c2"></a>

## Next pages — qradar_receiver.batch.max_bytes_disabled / a5c211f0cc79 / 4

- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-494d73db92f553ebadb34dbc35913c4b785e42cd0cac3b96fbe67393f6eb0524"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34ef6d3e4043796efb2b7d5551185f85db5638437574b0df4cc5dd44c15d8e9a"></a>

## qradar_receiver.batch.max_events_disabled — qradar_receiver.batch.max_events_disabled / 173689f66015 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61)
- qradar_receiver.batch.max_events_disabled

<a id="canonical-d10152deb0ec74139d2b01578eec86f350f8e2f953cf9850d680551f30cb79ed"></a>

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

<a id="canonical-805cae8edb4d9192d542c2f90411c21262d3de1f39886dbc9cdb168cafa2e18b"></a>

## Direct properties — qradar_receiver.batch.max_events_disabled / 173689f66015 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-008075aae5c6dd43ebf7596a4e15b7842ae14916bfeb9a3b5a0c8724496d20e1"></a>

## Next pages — qradar_receiver.batch.max_events_disabled / 173689f66015 / 4

- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-45bc78c0697d2d94c38d472c9ca67834d55f8162e1810df42a877271b6c165ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e651a5c28e73b82aae87df66f0536242c3d32311336542dd5364456534442dd4"></a>

## qradar_receiver.batch.timeout_seconds_default — qradar_receiver.batch.timeout_seconds_default / b97aac083db0 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61)
- qradar_receiver.batch.timeout_seconds_default

<a id="canonical-87570d7c5eadacfce5890c91f27d7211a89e6588ab86f012f823c0a8114df7b7"></a>

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

<a id="canonical-2d68e8928a75e88ce887a6dab09e9977561ca1bf758106a3ddb7597cf98f0797"></a>

## Direct properties — qradar_receiver.batch.timeout_seconds_default / b97aac083db0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-61f56455ef51f5ec64c57bccc18a40d2f84b2c1dd2ba44e97d9e32998820980d"></a>

## Next pages — qradar_receiver.batch.timeout_seconds_default / b97aac083db0 / 4

- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-fa49ced353fa20a904491b6c5ab2b2b093499dcf61f31d7975c3a6a5b4ca2e61)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f6bea6a4862ecfdcac771508f2c4b57d8cce6e7cbb3d9c10233b0995eda3d9d"></a>

## qradar_receiver.compression — qradar_receiver.compression / 9e27b474dd85 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- qradar_receiver.compression

<a id="canonical-333b8d781a27bb4f2ca2436d507ddb1fbe6473e28b5c4c300067c5710e2d5f82"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-6425f47ab203af5684af1b2e301c8bbacdde7287ec85a902188e89093008351b"></a>

## Direct properties — qradar_receiver.compression / 9e27b474dd85 / 3

- [compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-80e18658165226b07c77ee621f4783847548bc6774bce65a4fdd5798c8f181c3): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-89a5bc38e76d8b676fe87c851bd5cdf3cac64131d79ccacedfaa32618ca20515): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-cbc85b5788d7d70b6595890d0054572e630abf898e403b90fbc75122fe6a7c13): complete subsection reference.

<a id="canonical-80c3e2df948fc88c088217f88164d3d081ffb0ba84f52126d5f5bc84767dc841"></a>

## Next pages — qradar_receiver.compression / 9e27b474dd85 / 4

- [qradar_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-80e18658165226b07c77ee621f4783847548bc6774bce65a4fdd5798c8f181c3)
- [qradar_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-89a5bc38e76d8b676fe87c851bd5cdf3cac64131d79ccacedfaa32618ca20515)
- [qradar_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-cbc85b5788d7d70b6595890d0054572e630abf898e403b90fbc75122fe6a7c13)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-80e18658165226b07c77ee621f4783847548bc6774bce65a4fdd5798c8f181c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b22f552b11e5358b73e9536dcf9ee49e90c5e5dbed7c75da77fd4e9e8e059d5e"></a>

## qradar_receiver.compression.compression_default — qradar_receiver.compression.compression_default / 1841442ac63d / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4)
- qradar_receiver.compression.compression_default

<a id="canonical-4f5d489c596a20bac41f9f817b75f3adf68d092297cfe5b552e96a6b019d9a83"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-462948b41d333cd24d7c339912e3126b2b9c52a3848a225f67082e731ecbb648"></a>

## Direct properties — qradar_receiver.compression.compression_default / 1841442ac63d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4a0fe0ce31714d7a3f90eb6e9ec845662d8ed270606e842e17888c8d1e438ab"></a>

## Next pages — qradar_receiver.compression.compression_default / 1841442ac63d / 4

- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-89a5bc38e76d8b676fe87c851bd5cdf3cac64131d79ccacedfaa32618ca20515"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5c76831401b72e4800d10cd8c69ae56cf08d1a3607a72dc8c74305947321828"></a>

## qradar_receiver.compression.compression_gzip — qradar_receiver.compression.compression_gzip / 1054d8dea0a1 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4)
- qradar_receiver.compression.compression_gzip

<a id="canonical-efe622676ccb3473753dae8b15cabb2acf751fd60b422d47cdda9f1a03790ca4"></a>

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

<a id="canonical-0e0ef98c7d799db2ac313ae697409c230065f701cfd3ed2f24a71778ef9bfcb0"></a>

## Direct properties — qradar_receiver.compression.compression_gzip / 1054d8dea0a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-970285506c6e694a20c0a804866f5e4044473595a9d628d62faf71529d8aab4b"></a>

## Next pages — qradar_receiver.compression.compression_gzip / 1054d8dea0a1 / 4

- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-cbc85b5788d7d70b6595890d0054572e630abf898e403b90fbc75122fe6a7c13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86654daae1617e3c94e53cd684b012c208e50fc3b2111253601b3e82b4acde8c"></a>

## qradar_receiver.compression.compression_none — qradar_receiver.compression.compression_none / db6f8450921b / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4)
- qradar_receiver.compression.compression_none

<a id="canonical-4c8c45721e36189f9ccb2d754b1e28078dd899fecbd3e64a612910d1d6fabccb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-316f274d5fac819a7108d97f7ad935a781b1c6074971c5998caee68d3464633e"></a>

## Direct properties — qradar_receiver.compression.compression_none / db6f8450921b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6e1ddb3ad635b04bdd8544bec217500366fc5d69dd253eedc88c5b66a17345da"></a>

## Next pages — qradar_receiver.compression.compression_none / db6f8450921b / 4

- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1f9c7ab2cd063df89f058d0b6a2ae8cec41c49dd967396f6de6cc5879ce6fca4)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-55ec99baa6f49fbf244a8fbfdba9501b9d3552dc096a2d9bf1d6257c2613a3a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bb16420429b72a40a9ac53280bfef55241f82e7f504b494e5316684ae6a6b1f"></a>

## qradar_receiver.no_tls — qradar_receiver.no_tls / 75a9864abab7 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- qradar_receiver.no_tls

<a id="canonical-75a44f1d1031a8471baa0fd8a92ac2fe1f6c495e30fbdd73b3127432703625aa"></a>

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

<a id="canonical-c0fc4caaad2397e97d7b4d435def6aae14f256be3ca6a5948f9d593b3aba12c3"></a>

## Direct properties — qradar_receiver.no_tls / 75a9864abab7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be1ef322f1471c82b212753ba26184fa93ad9ea45ac135ef408df92ba8dcd604"></a>

## Next pages — qradar_receiver.no_tls / 75a9864abab7 / 4

- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-134256056864fa1a02534f338cf882c5ad5316a542794857df0a8ddf0098fefc"></a>

## qradar_receiver.use_tls — qradar_receiver.use_tls / cffe3dca4552 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- qradar_receiver.use_tls

<a id="canonical-aca335a47ecfe2fb33c34eab49101bc1a1029749d745bb8e64a4cbf13b287a72"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

<a id="canonical-6d5b0bdcfe2a26aca66e49f7538b77db9feabf7a2c7e6be712be3c89fac491c2"></a>

## Direct properties — qradar_receiver.use_tls / cffe3dca4552 / 3

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-fac72585cf67efc620f7373eaa6a91111fbdb3b1da04baffc61f5c43b50bbb57): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-cf7e3330382b856cd94c2ada9a059a30f5ce6525f846d03aa09f511e63135ec0): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-1b6ef19a97c7cae29923dae1a6e7f06c47f24e8fc616bc4b51a58ef45942e5b8): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-4830cf5602238781f99f58e59d1144988ca8e80a56db34fb2579a2fc82e31e12): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-ccae72ae8709e8c46ec76042c60fc19c7ed0dffc37b985e5346af762be202e17): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c967e66c3ca78e06dfb065213331847c2ca04af9f014dcea25c1d0f01d089bd7): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-795c32cedffd28a0fc7e0221a4b413d357c43cc33256809536d1fd168f30a0ba): complete subsection reference.

<a id="canonical-fb06c0d68e5a89ccab5a28718a3e2307d38b62212265f8244b2e8346d0a21732"></a>

<a id="canonical-ac698d7d4c375f94ed7331fca21de222a11a3ff4d6ac54ee6dccc41d8c5208da"></a>

## trusted_ca_url property — qradar_receiver.use_tls / cffe3dca4552 / 4

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2434ffd0df068b239e4e166aa63630003b1575742b450961dc78d50086977875"></a>

## Next pages — qradar_receiver.use_tls / cffe3dca4552 / 5

- [qradar_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-fac72585cf67efc620f7373eaa6a91111fbdb3b1da04baffc61f5c43b50bbb57)
- [qradar_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-cf7e3330382b856cd94c2ada9a059a30f5ce6525f846d03aa09f511e63135ec0)
- [qradar_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-1b6ef19a97c7cae29923dae1a6e7f06c47f24e8fc616bc4b51a58ef45942e5b8)
- [qradar_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-4830cf5602238781f99f58e59d1144988ca8e80a56db34fb2579a2fc82e31e12)
- [qradar_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-ccae72ae8709e8c46ec76042c60fc19c7ed0dffc37b985e5346af762be202e17)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c967e66c3ca78e06dfb065213331847c2ca04af9f014dcea25c1d0f01d089bd7)
- [qradar_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-795c32cedffd28a0fc7e0221a4b413d357c43cc33256809536d1fd168f30a0ba)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-fac72585cf67efc620f7373eaa6a91111fbdb3b1da04baffc61f5c43b50bbb57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-429fa943619d3093aa649f21986138a5e2ca6569494ae4beccd193133bced2d0"></a>

## qradar_receiver.use_tls.disable_verify_certificate — qradar_receiver.use_tls.disable_verify_certificate / b0632bef3277 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- qradar_receiver.use_tls.disable_verify_certificate

<a id="canonical-44514df5ca5709d389e614cd04a33f01d8f264e9f5abef71b706b91fd25a5ed8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-c21c5d8e476077cd5105289873c0fe0067a83c8f4120dda1f694d85ac1c078ce"></a>

## Direct properties — qradar_receiver.use_tls.disable_verify_certificate / b0632bef3277 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f21db1bd2fb7ac4dc6b2fc068514b72cf401aa29b731b14467ad2d98a67b2914"></a>

## Next pages — qradar_receiver.use_tls.disable_verify_certificate / b0632bef3277 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-cf7e3330382b856cd94c2ada9a059a30f5ce6525f846d03aa09f511e63135ec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecf5c8f13d0dce9c36a29b6cc2110cedb5a32468358e50ada89f85721ae2c117"></a>

## qradar_receiver.use_tls.disable_verify_hostname — qradar_receiver.use_tls.disable_verify_hostname / 9f81661143fb / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- qradar_receiver.use_tls.disable_verify_hostname

<a id="canonical-0f39c62fab243475d1d23aa99084782cb61357585962b55264df06512a271c61"></a>

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

<a id="canonical-23b3d727ce79a30d518758db4befdf23fb4ea6313cf65331025a714583b60c7b"></a>

## Direct properties — qradar_receiver.use_tls.disable_verify_hostname / 9f81661143fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7cb06ec09610047b9d2c2959d4af35354edc53383c322d0fc05819596a64f75"></a>

## Next pages — qradar_receiver.use_tls.disable_verify_hostname / 9f81661143fb / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-1b6ef19a97c7cae29923dae1a6e7f06c47f24e8fc616bc4b51a58ef45942e5b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2649a465b6066f71a54c73ca6ef712cb3d389ee7f2a39768e939308ad785116f"></a>

## qradar_receiver.use_tls.enable_verify_certificate — qradar_receiver.use_tls.enable_verify_certificate / 9d131c715285 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- qradar_receiver.use_tls.enable_verify_certificate

<a id="canonical-d3db25716d78ef3fc901f11e7ebbedfb653e6ec6244877b8e1240ea8e6f5f4c3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-acce40105acb1a2a698e28a7e02f0544fe33773ba69027a068f59ef87e818a3a"></a>

## Direct properties — qradar_receiver.use_tls.enable_verify_certificate / 9d131c715285 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a24dfbf8f8452be09eea682bf47777d1907466bd5e08efae509ca918b3dc8395"></a>

## Next pages — qradar_receiver.use_tls.enable_verify_certificate / 9d131c715285 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-4830cf5602238781f99f58e59d1144988ca8e80a56db34fb2579a2fc82e31e12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf40dd83a0525ac65915edb3ea78a8b213ae7cdb5f951c0c37461458f3f45997"></a>

## qradar_receiver.use_tls.enable_verify_hostname — qradar_receiver.use_tls.enable_verify_hostname / e4fc7703205a / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- qradar_receiver.use_tls.enable_verify_hostname

<a id="canonical-be0cf896738c5e38bd1612313a4c1a0f3212e9f26b3926ff03dd0073b30597c3"></a>

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

<a id="canonical-fe25917816d6a581e1df42379cb3d71cc3856978642c43f54dba903db03518f8"></a>

## Direct properties — qradar_receiver.use_tls.enable_verify_hostname / e4fc7703205a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3200adf0c65a6c0ea3175f9cab05fd62d42996f1a6668759300f2b4560e2f56"></a>

## Next pages — qradar_receiver.use_tls.enable_verify_hostname / e4fc7703205a / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-ccae72ae8709e8c46ec76042c60fc19c7ed0dffc37b985e5346af762be202e17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d898c0bca05fc67f28c0e8f6194b5279b03b3b50bdc0de308581b36aa3808227"></a>

## qradar_receiver.use_tls.mtls_disabled — qradar_receiver.use_tls.mtls_disabled / 4c0c166b3648 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- qradar_receiver.use_tls.mtls_disabled

<a id="canonical-4d09553a085ec5fe3b9732a823327c4e0339d5c73050e91cb8d1a7d96b9f178c"></a>

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

<a id="canonical-dace75a712dbfd4f02ce4fae5ee646ddd1652f8d6bacea2ef0ebc69804d7e9f3"></a>

## Direct properties — qradar_receiver.use_tls.mtls_disabled / 4c0c166b3648 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9c681701eba85bf4cfe31846e3c134c3208aa85397f4b704bda0fdd16dcd231"></a>

## Next pages — qradar_receiver.use_tls.mtls_disabled / 4c0c166b3648 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-c967e66c3ca78e06dfb065213331847c2ca04af9f014dcea25c1d0f01d089bd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42fd17f46ab9358fa838e708c82ab47ee691f5b3659b5f8f4d26cba1f4506f42"></a>

## qradar_receiver.use_tls.mtls_enable — qradar_receiver.use_tls.mtls_enable / 8ad8581979ed / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- qradar_receiver.use_tls.mtls_enable

<a id="canonical-5203b1edc5f8da69ebea5e9cdff7bad9eecd3bd33da6a7566b4729d5e9c67068"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

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

<a id="canonical-fc42afccf78c765b0aa98318dd9e07cf3a59c0964d55ebba668249ced7609ffc"></a>

## Direct properties — qradar_receiver.use_tls.mtls_enable / 8ad8581979ed / 3

<a id="canonical-4bb979d800e90d819b4f0ef572613be12191da15ae8865b39a8ce33443d1fb7a"></a>

<a id="canonical-73866af393d11971abda341077e2f7c51a6f125b011ca7e3a0b1309bb46d3bb8"></a>

## certificate property — qradar_receiver.use_tls.mtls_enable / 8ad8581979ed / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-d0c6c715ea45d38fd6d5996b76301ba30ac04e6a129e1cb996c13974912d2c72): complete subsection reference.

<a id="canonical-979d088536bf806fb8668a00449d3c2d2a12e8bdaeda5be95c3c7507da9b6a8f"></a>

## Next pages — qradar_receiver.use_tls.mtls_enable / 8ad8581979ed / 5

- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-d0c6c715ea45d38fd6d5996b76301ba30ac04e6a129e1cb996c13974912d2c72)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-d0c6c715ea45d38fd6d5996b76301ba30ac04e6a129e1cb996c13974912d2c72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d7d7d0de7cb7c70f753f20abe8d0b7e49c4170357d6f5c2cc1c421c42d5f188"></a>

## qradar_receiver.use_tls.mtls_enable.key_url — qradar_receiver.use_tls.mtls_enable.key_url / c1364beb4cdd / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-71ebb33f827d1c6f274624cc4b12b2f74df3dfba6681014794fca1fc58ee64fe)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-b92792beb6b0768e505b489f160b029bb8d833695a382acbf5b46a70680d13a3)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-a588b567a94fa50cbd3de89986489e164caf88540014417f0504de59af2171b5)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c967e66c3ca78e06dfb065213331847c2ca04af9f014dcea25c1d0f01d089bd7)
- qradar_receiver.use_tls.mtls_enable.key_url

<a id="canonical-c7aa5b3930e519b01d85c4e114d15eb07a231e8102c1c51e465aaa05427942c7"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-fd796eb0df09536a3adb0b094530c812d3bfb9eff70c4810e8325ce17d33f30d"></a>

## Direct properties — qradar_receiver.use_tls.mtls_enable.key_url / c1364beb4cdd / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-e042091722810fd12441a50c232b2ece6d58fb07cb082343224b1688c9eebd95): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-7e3cb6f28d0dd357bbb36d620bd9e5d500b3984fe74f976f4e40ac289dc48e68): complete subsection reference.

<a id="canonical-a0b8eeac8b13a485d07e2e3214639249b5186fcdd937ac8a8eff894aea48c73d"></a>

## Next pages — qradar_receiver.use_tls.mtls_enable.key_url / c1364beb4cdd / 4

- [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-e042091722810fd12441a50c232b2ece6d58fb07cb082343224b1688c9eebd95)
- [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-7e3cb6f28d0dd357bbb36d620bd9e5d500b3984fe74f976f4e40ac289dc48e68)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-c967e66c3ca78e06dfb065213331847c2ca04af9f014dcea25c1d0f01d089bd7)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-e042091722810fd12441a50c232b2ece6d58fb07cb082343224b1688c9eebd95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
