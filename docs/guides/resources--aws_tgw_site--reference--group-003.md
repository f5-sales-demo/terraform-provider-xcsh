---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-a2d241a1a9c6de9111f8ec182a8ca2167e76eba2ec6ec30e926e154c519cb820"></a>

## vn_config.allowed_vip_port — vn_config.allowed_vip_port / 924feab6554e / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.allowed_vip_port

<a id="canonical-3a6f8b9cae264f949303180413bbf7a52575cdd9ebc8c79c1d9ddb01f9f0c927"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-773805299b321b024d9cf5c6416486cda01ff522353d6cb61019658e415c5c68"></a>

## Direct properties — vn_config.allowed_vip_port / 924feab6554e / 3

- [custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-e9757e7ed2ed675cf0f1abbdec87e110d2e08b5d40be1e78ceb0a213d6972366): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-7b75b6b4f0bbd53077476f54082eb2baca00ec2e4c06a644fed6434150d60dfe): complete subsection reference.

- [use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-b16ad48663791b4fe48ab25f64f764b0955e8f05b011be48c09445dc315a38ed): complete subsection reference.

- [use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-dfd62a3c2fafb43495ec70b8929f5820c2b7b42cc52df1dde6c8df0c7c3eb6e0): complete subsection reference.

- [use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-bcb760c9fa2176028b1b7cd59de698ba51937c8a389c38042734ec0c5942c61a): complete subsection reference.

<a id="canonical-ab696389273454521167d2ef27196829b51d3becb9ba9c5170f768b2969c777b"></a>

## Next pages — vn_config.allowed_vip_port / 924feab6554e / 4

- [vn_config.allowed_vip_port.custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-e9757e7ed2ed675cf0f1abbdec87e110d2e08b5d40be1e78ceb0a213d6972366)
- [vn_config.allowed_vip_port.disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-7b75b6b4f0bbd53077476f54082eb2baca00ec2e4c06a644fed6434150d60dfe)
- [vn_config.allowed_vip_port.use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-b16ad48663791b4fe48ab25f64f764b0955e8f05b011be48c09445dc315a38ed)
- [vn_config.allowed_vip_port.use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-dfd62a3c2fafb43495ec70b8929f5820c2b7b42cc52df1dde6c8df0c7c3eb6e0)
- [vn_config.allowed_vip_port.use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-bcb760c9fa2176028b1b7cd59de698ba51937c8a389c38042734ec0c5942c61a)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-e9757e7ed2ed675cf0f1abbdec87e110d2e08b5d40be1e78ceb0a213d6972366"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc4818f5e3cd90cd34f4872fb0272d6b31666e9c08db413f73ef5d75d7f94fc3"></a>

## vn_config.allowed_vip_port.custom_ports — vn_config.allowed_vip_port.custom_ports / 57986d9a82ac / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- vn_config.allowed_vip_port.custom_ports

<a id="canonical-a4a61d6c084e6d4861b8dacf869d4a92901320edb1e37b1c0fd7a086bbd6a455"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
```

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

Terraform syntax:

```terraform
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-53796b5c34d0736a85815208a9f5c2fbc1d346a6cabb02afba1f8d64a0d024b2"></a>

## Direct properties — vn_config.allowed_vip_port.custom_ports / 57986d9a82ac / 3

<a id="canonical-0ad6517e5187bbdac18c59b1cb66d3d2431abc45f4f2fc3d136b9f105423be0a"></a>

<a id="canonical-c92c97ee504f917695ac21a457c684be2c036102385e47d21f4e216380d7d00a"></a>

## port_ranges property — vn_config.allowed_vip_port.custom_ports / 57986d9a82ac / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-29ee331c6a945581bfb4c46f79fd219e8138b551195e125c957c70fbc7324a7e"></a>

## Next pages — vn_config.allowed_vip_port.custom_ports / 57986d9a82ac / 5

- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-7b75b6b4f0bbd53077476f54082eb2baca00ec2e4c06a644fed6434150d60dfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8688a8c1837f2bd9edcb69e81df1cf37f4d69b603db7744c299d7d73ee364869"></a>

## vn_config.allowed_vip_port.disable_allowed_vip_port — vn_config.allowed_vip_port.disable_allowed_vip_port / e8d15eb38000 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- vn_config.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-ca6cc2ac5593cd01cc77981d09bd43cd805c49dc8eccfcdfaedbbd21bd48f762"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_allowed_vip_port = {}
```

<a id="canonical-80818adc13a744d2a8328897d622d861d10906db5afeda4552add3466aa064a8"></a>

## Direct properties — vn_config.allowed_vip_port.disable_allowed_vip_port / e8d15eb38000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f8f428e38f5a49168226542f622078866a88e42003fa79245b1bbabed3db016a"></a>

## Next pages — vn_config.allowed_vip_port.disable_allowed_vip_port / e8d15eb38000 / 4

- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b16ad48663791b4fe48ab25f64f764b0955e8f05b011be48c09445dc315a38ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd1c9fde7c3040760179fab649df5fe64ca50820e3ffeca40877ce5939d2093e"></a>

## vn_config.allowed_vip_port.use_http_https_port — vn_config.allowed_vip_port.use_http_https_port / 2d56926500e9 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- vn_config.allowed_vip_port.use_http_https_port

<a id="canonical-ca2d97ad209edeb8d550b7d897195e30dd36efa8b9b275f425ab079b0f61ef4f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_http_https_port = {}
```

<a id="canonical-e8dc157bb5065e6878c35cdc374b858571cc301c87b42d6d5c9fe8ba1ce94743"></a>

## Direct properties — vn_config.allowed_vip_port.use_http_https_port / 2d56926500e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-737dd79765c0100f1fbca3a8bd80bf6aa05637eddec1bbb1d34e73782644d773"></a>

## Next pages — vn_config.allowed_vip_port.use_http_https_port / 2d56926500e9 / 4

- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-dfd62a3c2fafb43495ec70b8929f5820c2b7b42cc52df1dde6c8df0c7c3eb6e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6cc64acf2ca6a4d618c6a03e399f941d35cf965e2565e31d7c15c01bde4963a"></a>

## vn_config.allowed_vip_port.use_http_port — vn_config.allowed_vip_port.use_http_port / 40a93223dcc1 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- vn_config.allowed_vip_port.use_http_port

<a id="canonical-613a685dac6043a8f4f6ab107b38b57e4348dee7b6258b628215cd6aba386d2b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_http_port = {}
```

<a id="canonical-5890d2a100eb8922bb658e2c31070a1baf15b4ba1979239292ef4014620c5e10"></a>

## Direct properties — vn_config.allowed_vip_port.use_http_port / 40a93223dcc1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-998e1d420e6e3766b09f4daeb6deeea2a9b0825c62c3d0ea6c750d09ca3979b4"></a>

## Next pages — vn_config.allowed_vip_port.use_http_port / 40a93223dcc1 / 4

- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-bcb760c9fa2176028b1b7cd59de698ba51937c8a389c38042734ec0c5942c61a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c415029527729f0548da010d20125e89fdf0acd08f24ce668fc66942c22fa1e4"></a>

## vn_config.allowed_vip_port.use_https_port — vn_config.allowed_vip_port.use_https_port / 7aede36d9603 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- vn_config.allowed_vip_port.use_https_port

<a id="canonical-2ca78f0ff3515f2b3a1f5cd5ed2bcc2c4f4e4542bc11b04b893ba6e32f65faf2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_https_port = {}
```

<a id="canonical-46945a8fd73c7a2e48605b3b435164aef14a3737b21cf7c5200656a16c55bab8"></a>

## Direct properties — vn_config.allowed_vip_port.use_https_port / 7aede36d9603 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-675f20e819a3e9c520a8f583f22a4d9f10ca7bf802010314e40c36baa9f4edf7"></a>

## Next pages — vn_config.allowed_vip_port.use_https_port / 7aede36d9603 / 4

- [vn_config.allowed_vip_port](resources--aws_tgw_site--reference--group-002.md#canonical-c3484a7e11aa54ebb63153fea606735eb17cf3b9a1d09f03216ff561b0fa5c27)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d2fb0b5091cfef7d24b8357bb45fadf787c142cec726d92f4a54fbea2392b78"></a>

## vn_config.allowed_vip_port_sli — vn_config.allowed_vip_port_sli / 3a7f98462899 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.allowed_vip_port_sli

<a id="canonical-ad51201dfcdfc13578233e5cde7d14f603c24abb79768cfdf2812a4312c849b4"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-f2f10952a76b9ccd4b972e4b5cbe75ba87f3f3a26aa7fed22f5204120f2ce924"></a>

## Direct properties — vn_config.allowed_vip_port_sli / 3a7f98462899 / 3

- [custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-f9d2d155469fa101b22d6b35e16bfde7355b6b1959d0ab5f3e182ea15d08ef7d): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-05ef7f91a4b903e011d776dd2fe1cef81bec42724fef3b9a8d03079b33bf6742): complete subsection reference.

- [use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-9e2ed74d2b4ba78c0d7653986d21070260c28733f1bf3861cee2c6ceb5072d91): complete subsection reference.

- [use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-4d4e1f63ef3a495f45938999680e53f68ffbb68b7c5b512866147fa376e9fc9b): complete subsection reference.

- [use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-f64b41ba341ef67361df5dfe6d7c7cdf269e9d04eaa235b670ce0db3ba625799): complete subsection reference.

<a id="canonical-153676528780b44d12d45cdb3272a5a737aa388387ba47d0a5a7a0b9928d5712"></a>

## Next pages — vn_config.allowed_vip_port_sli / 3a7f98462899 / 4

- [vn_config.allowed_vip_port_sli.custom_ports](resources--aws_tgw_site--reference--group-003.md#canonical-f9d2d155469fa101b22d6b35e16bfde7355b6b1959d0ab5f3e182ea15d08ef7d)
- [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_tgw_site--reference--group-003.md#canonical-05ef7f91a4b903e011d776dd2fe1cef81bec42724fef3b9a8d03079b33bf6742)
- [vn_config.allowed_vip_port_sli.use_http_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-9e2ed74d2b4ba78c0d7653986d21070260c28733f1bf3861cee2c6ceb5072d91)
- [vn_config.allowed_vip_port_sli.use_http_port](resources--aws_tgw_site--reference--group-003.md#canonical-4d4e1f63ef3a495f45938999680e53f68ffbb68b7c5b512866147fa376e9fc9b)
- [vn_config.allowed_vip_port_sli.use_https_port](resources--aws_tgw_site--reference--group-003.md#canonical-f64b41ba341ef67361df5dfe6d7c7cdf269e9d04eaa235b670ce0db3ba625799)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-f9d2d155469fa101b22d6b35e16bfde7355b6b1959d0ab5f3e182ea15d08ef7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da499322f4d88c6afc01bbcc677ecb44e2d29c3599f2a9fee68c6292257c3a4f"></a>

## vn_config.allowed_vip_port_sli.custom_ports — vn_config.allowed_vip_port_sli.custom_ports / 37542b53ebf5 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- vn_config.allowed_vip_port_sli.custom_ports

<a id="canonical-4b1d99a9291c2d1ec427d94b06de02845bd1299ade1186d0fb26341a3096c14c"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
```

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

Terraform syntax:

```terraform
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-9748220291889f12b49ab4c9dc6422a2702db96b1f7d3c30a38f34ccbf5d2a3f"></a>

## Direct properties — vn_config.allowed_vip_port_sli.custom_ports / 37542b53ebf5 / 3

<a id="canonical-0dd27f44b80cbd3f01cd429789897d3931878363e88a8bd2b215ee8e969028c9"></a>

<a id="canonical-c6be8ce739a11fe2cb6fc32f08538cfdf95fe8598c3de5aa8544a5a66fbb5c58"></a>

## port_ranges property — vn_config.allowed_vip_port_sli.custom_ports / 37542b53ebf5 / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-8b22592ba940a5169b21395d893cc77371a18bc70762be68c30de40eaacdb0bf"></a>

## Next pages — vn_config.allowed_vip_port_sli.custom_ports / 37542b53ebf5 / 5

- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-05ef7f91a4b903e011d776dd2fe1cef81bec42724fef3b9a8d03079b33bf6742"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aec8606156575af5b661be0443ee27c95315fb9c85513df7f04dcc2dfe5af62"></a>

## vn_config.allowed_vip_port_sli.disable_allowed_vip_port — vn_config.allowed_vip_port_sli.disable_allowed_vip_port / 2c49419bb443 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- vn_config.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-2a553c75361f0add0104d8602b968fadd30eb0ed204339022383847cfc156b8a"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_allowed_vip_port = {}
```

<a id="canonical-12e8e9c273f89b0dfaa6cc7c62cbc885d7c173491e4553e155ba907730c191c8"></a>

## Direct properties — vn_config.allowed_vip_port_sli.disable_allowed_vip_port / 2c49419bb443 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a878f1a0893a231ac53a2a72e6d3cf890111d7bd2acf54460b9b1792fb70275"></a>

## Next pages — vn_config.allowed_vip_port_sli.disable_allowed_vip_port / 2c49419bb443 / 4

- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-9e2ed74d2b4ba78c0d7653986d21070260c28733f1bf3861cee2c6ceb5072d91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de11685b6e6f685909bf31ad8c153d3aa63ab03be6836be29da13b64fd86d289"></a>

## vn_config.allowed_vip_port_sli.use_http_https_port — vn_config.allowed_vip_port_sli.use_http_https_port / b503aaa90cc3 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- vn_config.allowed_vip_port_sli.use_http_https_port

<a id="canonical-5b4faa612541843bf9e4df1263f8669f12e2497b815fadfdb2ae0de6e5bad6d0"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_http_https_port = {}
```

<a id="canonical-558a519ddc383abf85c31d9d226bcebc26b86f102b783d668dcd5cd3ec6fb0ed"></a>

## Direct properties — vn_config.allowed_vip_port_sli.use_http_https_port / b503aaa90cc3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da737811f376ef65738d6cce2bbbd981857e4bc442d9f1c8dd9392190ab78f66"></a>

## Next pages — vn_config.allowed_vip_port_sli.use_http_https_port / b503aaa90cc3 / 4

- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-4d4e1f63ef3a495f45938999680e53f68ffbb68b7c5b512866147fa376e9fc9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ae85b98ac8491d7f77ef7bb8b0949401c887e7039bc865170c263166d8e4781"></a>

## vn_config.allowed_vip_port_sli.use_http_port — vn_config.allowed_vip_port_sli.use_http_port / 32c93d81174c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- vn_config.allowed_vip_port_sli.use_http_port

<a id="canonical-8c4aedd707bb0c993bc113bf69942b21e18e89e9b060fa4b43b48ef87635e1ea"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_http_port = {}
```

<a id="canonical-e9898347963344e410882e3b607fa2f518fd2ba279888d2f8b2f83d19bfbe61a"></a>

## Direct properties — vn_config.allowed_vip_port_sli.use_http_port / 32c93d81174c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ea288ab48c9486795a9f29106fd67e4a68c239c3b087761a03ffdc13f21bc5c"></a>

## Next pages — vn_config.allowed_vip_port_sli.use_http_port / 32c93d81174c / 4

- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-f64b41ba341ef67361df5dfe6d7c7cdf269e9d04eaa235b670ce0db3ba625799"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb5c21600b2588ed296fb98c38f6e36d3bdd969796b0883c3351019c281613a8"></a>

## vn_config.allowed_vip_port_sli.use_https_port — vn_config.allowed_vip_port_sli.use_https_port / d2bf8842a9be / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- vn_config.allowed_vip_port_sli.use_https_port

<a id="canonical-cbb22e0d5930a60ccb875d5e6955225e0014a779e1c4b6d2467275b1321552a1"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_https_port = {}
```

<a id="canonical-4381ab9f8540b2d8dba1d0ebf8e76dda096a17237a8b93bce42dd445db369a72"></a>

## Direct properties — vn_config.allowed_vip_port_sli.use_https_port / d2bf8842a9be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-047feb442939aa93902b7ca4f22b42125970c910c72f92f5e8347f1f3bb11224"></a>

## Next pages — vn_config.allowed_vip_port_sli.use_https_port / d2bf8842a9be / 4

- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--reference--group-003.md#canonical-71cf965bb9aee847d2f192259885d373c912551dbdbdf6f48b9964f8776ec329)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-fb609dc791b3d1048ce594023281cdbcf0a207f1656de710ea962117ede022be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-771c620fbe6b8de56ae701bebd6de53741ea1fc27eea3e15b2f48413f558f6c6"></a>

## vn_config.dc_cluster_group_inside_vn — vn_config.dc_cluster_group_inside_vn / 4739091dc480 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.dc_cluster_group_inside_vn

<a id="canonical-32b5d510ca4d7dd340a778452f046ebae147467f1eaeb52fc8b42251ae4f7953"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
dc_cluster_group_inside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-c43ce33d0e5d35ef9f466e1e9d69b3295112df7ec64f7452e011c6e9a98c9fe8"></a>

## Direct properties — vn_config.dc_cluster_group_inside_vn / 4739091dc480 / 3

<a id="canonical-33a76fbf3b2176dff8855a7abe1503cd32af824e2d543d1ea1475b8f9c98936f"></a>

<a id="canonical-3666c6f9a6f2f618414a1d7f5e16d558593c5c1e00aad89d759be7e79d3e2b67"></a>

## name property — vn_config.dc_cluster_group_inside_vn / 4739091dc480 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-5aea276ec576d450ec1352ee0393948cb1cf60a57696fff9dfc09a75c431b66a"></a>

<a id="canonical-10fe91a32143e715c72befc0d6000e49b90761f4f2a3efe846b8bbaca5554a89"></a>

## namespace property — vn_config.dc_cluster_group_inside_vn / 4739091dc480 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0b73b1aabbf68a4f05b02d1baa79805eee6a8bb4e4233cc446b5818b5a9493f1"></a>

<a id="canonical-1bc23d96d07b4d8a57c0a9d0094670d262d319f1aac1299e9ab047836ef4c134"></a>

## tenant property — vn_config.dc_cluster_group_inside_vn / 4739091dc480 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-88ea25a73e1c899f15f88808e19ec3374cdce8096da7b94830edf46ed6308c0c"></a>

## Next pages — vn_config.dc_cluster_group_inside_vn / 4739091dc480 / 7

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-63828422b1f3043f756939e95d1857c4026f64a856a9f5b90e54537161daf8dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6222cae1785063f35c3cd791e25e85ae9df09c94434c288dd8ae162526489cd4"></a>

## vn_config.dc_cluster_group_outside_vn — vn_config.dc_cluster_group_outside_vn / a7dd835e51a4 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.dc_cluster_group_outside_vn

<a id="canonical-0b97b402fe8921d8338faa719c984a3c38e46efbd16c1e78662e20340f66ca50"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
dc_cluster_group_outside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-d44d52fd83fa4c6d8bbe63889fbb0a49bdbb2e08173b562384c7e9e2868bf81b"></a>

## Direct properties — vn_config.dc_cluster_group_outside_vn / a7dd835e51a4 / 3

<a id="canonical-631a660b961c0a07c29866c641aaede62b75957b7c7cef385252ca4ce81e2633"></a>

<a id="canonical-c3d744e55d3c1071c73cdf5d93adc2b9bf27c2a6ad79d06591fbe1c81aefb139"></a>

## name property — vn_config.dc_cluster_group_outside_vn / a7dd835e51a4 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-20813f8aacf05df081eb1b48f584619ec980826dd15b71a5b33109e28025d9cc"></a>

<a id="canonical-d4fab4e835ef3b40a310b521fc23c5d20b15306f1b7005f555092c860443a346"></a>

## namespace property — vn_config.dc_cluster_group_outside_vn / a7dd835e51a4 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-daee89284ae1e9d08d0d825dbe7c7544944d5f375448001fc500bff13cb3306e"></a>

<a id="canonical-793ad8fd521e66fba3cf7c44b83bd4480288d35c808b1294ccedc896b1d3a165"></a>

## tenant property — vn_config.dc_cluster_group_outside_vn / a7dd835e51a4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-f8d005d3ca37b2e74f8655d41ce70db455eacdcb91f5a99687f18ac6d878f34d"></a>

## Next pages — vn_config.dc_cluster_group_outside_vn / a7dd835e51a4 / 7

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b53e1f4de5df514997708043c8890b26b689e1e2e71dcc52729681e8bfabfaad"></a>

## vn_config.global_network_list — vn_config.global_network_list / 00a4fa830bb1 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.global_network_list

<a id="canonical-9be859efac131b19296b21f17afe56fa2ee646558ed8e588e6da15241b9bf68c"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
```

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

Terraform syntax:

```terraform
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-cbaec75ba5af1604e0efb8349180eca614f7ad76289e789612b710cf6b771688"></a>

## Direct properties — vn_config.global_network_list / 00a4fa830bb1 / 3

- [global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e): complete subsection reference.

<a id="canonical-3333e126f9829fb2860ffaa3b56658480c64c804741b5c822c05abedf4f379d2"></a>

## Next pages — vn_config.global_network_list / 00a4fa830bb1 / 4

- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9ed71368e5afe48b6b6fb1e8ef4692b7f23896e3d700fe4226bc7065febf26d"></a>

## vn_config.global_network_list.global_network_connections — vn_config.global_network_list.global_network_connections / 2726426b9adb / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c)
- vn_config.global_network_list.global_network_connections

<a id="canonical-37ea0243159277f22c59c0c97a84c938f112aa88f7c72487d264cbe2c5b886aa"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-c842682845fdbc87215ad2b5fad265bdf86f2e67d686fd6853bd90df7fe2e604"></a>

## Direct properties — vn_config.global_network_list.global_network_connections / 2726426b9adb / 3

- [sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-7b31ec5741a0f7ea2a9c0648524dd81b07800fc474248c38ce402e588ce90874): complete subsection reference.

- [slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-3b22981252b34d5233498d04fa8f21ef802c694370aaafb277c400e188947ab3): complete subsection reference.

<a id="canonical-e8dff8bce6feb7d2148bde1d22ce04809792cef039cb994b1ea7285d169e43d3"></a>

## Next pages — vn_config.global_network_list.global_network_connections / 2726426b9adb / 4

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-7b31ec5741a0f7ea2a9c0648524dd81b07800fc474248c38ce402e588ce90874)
- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-3b22981252b34d5233498d04fa8f21ef802c694370aaafb277c400e188947ab3)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-7b31ec5741a0f7ea2a9c0648524dd81b07800fc474248c38ce402e588ce90874"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e8710f04470648466695c45d04f110cd9d8d3399134adcccfc1909a441dcd2d"></a>

## vn_config.global_network_list.global_network_connections.sli_to_global_dr — vn_config.global_network_list.global_network_connections.sli_to_global_dr / 3ccd3d0c40fa / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-45056b4c42ef5d1bed309f07f9249cffe13ae4396c1d24150d74c1f518241a4e"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-9ca68fd976889b28b116a5d266e6cec018ecaa66cf13f3d13f0307475d4017c3"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.sli_to_global_dr / 3ccd3d0c40fa / 3

- [global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-a1c391a938cb7cbb8d55a1a395b3821f42df11f187bdcde6c2c4e002eb52f4e7): complete subsection reference.

<a id="canonical-967324d4b602366d39c1e16be3030ae17bc6f1f2e6d766f8aa2c7d9642d3be58"></a>

## Next pages — vn_config.global_network_list.global_network_connections.sli_to_global_dr / 3ccd3d0c40fa / 4

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-a1c391a938cb7cbb8d55a1a395b3821f42df11f187bdcde6c2c4e002eb52f4e7)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-a1c391a938cb7cbb8d55a1a395b3821f42df11f187bdcde6c2c4e002eb52f4e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1959e943371511e13c9114f66dc29ec80d55e39083713ffaf55f020eae0460f"></a>

## vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 69e3319d8d8a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e)
- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-7b31ec5741a0f7ea2a9c0648524dd81b07800fc474248c38ce402e588ce90874)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-e52ca3b1482b677725071f34650551ffac08cda0acf90380db7acdc3f0ab12d0"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-f5daa3df08d1da31f6c4885ae693fc7919cf24a0968f48f73496de7b7bd3a215"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 69e3319d8d8a / 3

<a id="canonical-230fefb0eabe1fd4d7b640e5ea954a6a1f365c2a1c66bc6392dc2dbc8faf4593"></a>

<a id="canonical-bff6e04aa302477c4f18031d0191ca829bda710602a53be5da31e42f3edbb2cb"></a>

## name property — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 69e3319d8d8a / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-a1f80918bc704635371ead6c736b70c21018657a837c30ce15e479f02a85f60e"></a>

<a id="canonical-a69d3a7b9dd43c47262b9f694dd992afab64f9b937e879da49b7fc32526f5c3e"></a>

## namespace property — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 69e3319d8d8a / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-6660550b14502c1f18b78136ee16c5c7b4c8318bdbc926889d44bceb0195b282"></a>

<a id="canonical-4f5430e7c469a7954e208747ddee73b5ce591fa17e07046d34004dce51d19ae0"></a>

## tenant property — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 69e3319d8d8a / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-4ee8136c9f4a0b3190d976d4303c083573165e275a6c77bdcb25e8a5f327dd55"></a>

## Next pages — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 69e3319d8d8a / 7

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-7b31ec5741a0f7ea2a9c0648524dd81b07800fc474248c38ce402e588ce90874)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-3b22981252b34d5233498d04fa8f21ef802c694370aaafb277c400e188947ab3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e54c3de191e458c0b8a8402c1c083d5928d816123d8edd77f60607669e84990"></a>

## vn_config.global_network_list.global_network_connections.slo_to_global_dr — vn_config.global_network_list.global_network_connections.slo_to_global_dr / d84fbffcfed2 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-4b220696ec5e61951cba0de0b449596e66c75705eb5910e5b6664dabfbb3b53f"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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

Terraform syntax:

```terraform
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1c71df745e86af9609dfeb9e8cffc72896db546c0b7cd77b699dbebfe6cfa271"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.slo_to_global_dr / d84fbffcfed2 / 3

- [global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-062614525649fb59904cb50c0e1dc50a374eae59505a7ee3802afe17ec55b382): complete subsection reference.

<a id="canonical-22855523d1920cdbb92bdd18ecfec0affa83bfeaed6a452a97793cfa132fa15e"></a>

## Next pages — vn_config.global_network_list.global_network_connections.slo_to_global_dr / d84fbffcfed2 / 4

- [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_tgw_site--reference--group-003.md#canonical-062614525649fb59904cb50c0e1dc50a374eae59505a7ee3802afe17ec55b382)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-062614525649fb59904cb50c0e1dc50a374eae59505a7ee3802afe17ec55b382"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05d1791a4148cc52b3399c8314ad3bb5f5e76d26bb1b487e8aed890a1579efe1"></a>

## vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / a5c9c8ecc51e / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.global_network_list](resources--aws_tgw_site--reference--group-003.md#canonical-7501fd7d1f38e711b5a7cd500a3e399df66cb7a1dc4406d50771d2f3bbb5ec0c)
- [vn_config.global_network_list.global_network_connections](resources--aws_tgw_site--reference--group-003.md#canonical-c85400b92f06037de3425e2b445f4869d7bc5c65795ae0d0b80ca65089426d2e)
- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-3b22981252b34d5233498d04fa8f21ef802c694370aaafb277c400e188947ab3)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-0dbd6d4cc42a533d46891f925bb369b8e105d1a7bf9081c8083dc6fc99d7c4d2"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a44d6e48815955c53976066543d6184f065acf14957c87b2173d2680debaed6"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / a5c9c8ecc51e / 3

<a id="canonical-3d0d1c1c9e0557b70a1856e764b807e0cb10c23296fd74ffec02fcc685542b89"></a>

<a id="canonical-ce5be192e04cecd6576ca85ae803a0697ea8fae09a5111f0605a67667f7d4817"></a>

## name property — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / a5c9c8ecc51e / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-e735588c4b37f5d440c48cc3504c46265716e6f132fff7f1c0a71042fc367d59"></a>

<a id="canonical-6742549ee78e19ffeb09ae3eaf89d74f3b1127c905a5cad07d321912287060f3"></a>

## namespace property — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / a5c9c8ecc51e / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-34667d990cc7695a7d55bb87909d9c354cbc729d541eeca24bfd2a7b8071790b"></a>

<a id="canonical-77fdf7f4927ab70fc55955cbaa155c96988d43f2c5e7b7fc1d42d20448cf02d9"></a>

## tenant property — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / a5c9c8ecc51e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-41bc99d707fc064abe354c8b3dca0a52276728171d0e760dd0c719127d9413d1"></a>

## Next pages — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / a5c9c8ecc51e / 7

- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--reference--group-003.md#canonical-3b22981252b34d5233498d04fa8f21ef802c694370aaafb277c400e188947ab3)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9653b800cdeaeff3505e583d9e36d8f7ad46f26e4871c997aa2f2932346065ab"></a>

## vn_config.inside_static_routes — vn_config.inside_static_routes / c43f1b3aa6cc / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.inside_static_routes

<a id="canonical-e9d4c945d12a5181370ebf86c118bdd8e4f7fbc7d13111b80e88b55fca9eee7d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
```

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

Terraform syntax:

```terraform
inside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-101a143ad0b969fe38d9f29a294250cc196bcc84be7d30c7103466671810a2a1"></a>

## Direct properties — vn_config.inside_static_routes / c43f1b3aa6cc / 3

- [static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed): complete subsection reference.

<a id="canonical-e7ec9314baeef95d6b584a52c1a65aef2d203e710077f051b523090060c726ea"></a>

## Next pages — vn_config.inside_static_routes / c43f1b3aa6cc / 4

- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed50ef3acc3a22eac134e00ef6000ecd180a91d2726a8b754d0862decc1568a0"></a>

## vn_config.inside_static_routes.static_route_list — vn_config.inside_static_routes.static_route_list / b0af35fe542b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- vn_config.inside_static_routes.static_route_list

<a id="canonical-bf19c8f28b7dae1f59db06764aed1a0773128a91f6e449a46554812b06ca021e"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-7c64d71d4c40ea0ac8b7c9b4b417c81263db075ea721d8c88b0cd1f51c858d4b"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list / b0af35fe542b / 3

- [custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8): complete subsection reference.

<a id="canonical-f9c87c320f33f80c511b8cd297c5d1090e2bda0d656418e2e7d2efdedc556e6f"></a>

<a id="canonical-e3a061c7b6d314e80b8f9ddaf826539f2720c831267e1f29b5a757475bbe139b"></a>

## simple_static_route property — vn_config.inside_static_routes.static_route_list / b0af35fe542b / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2fbfc210ea9d3e1b7de60a8761a7ba29da19f3422c07124362057d77c57b55e5"></a>

## Next pages — vn_config.inside_static_routes.static_route_list / b0af35fe542b / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3c29643ac6b1a25dc1af391954168cfafd6aa8554b8666d79c9d0526a36d015"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route — vn_config.inside_static_routes.static_route_list.custom_static_route / 35bf2d2d4d21 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- vn_config.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-141186e7f666d11a3865c5cfc02f44fa04e72a2b46b8521306440bde02914f21"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
```

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

Terraform syntax:

```terraform
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-3c164119d9a463c44d1210deda459059c9893bce919140bda1aede7bb6770259"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route / 35bf2d2d4d21 / 3

<a id="canonical-572b299eac60cccc24853cdd3b6892a8ba2a515cbf6e1964b47a191eb5eccde0"></a>

<a id="canonical-eb6cde3850fe9f2555e3e9cb60a8d4921d5ab560f87ebc9dd3bdfc1fbdfdf3d6"></a>

## attrs property — vn_config.inside_static_routes.static_route_list.custom_static_route / 35bf2d2d4d21 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_tgw_site--reference--group-003.md#canonical-33ad9473556950b1d8eaeae633e15718a70e5415d43d3f778d28edc383a2e10a): complete subsection reference.

- [nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0): complete subsection reference.

- [subnets](resources--aws_tgw_site--reference--group-003.md#canonical-b649a7e018cabe1e1f1ed42021ffe62bcaf11f70355eb01c1b21b08646f7df04): complete subsection reference.

<a id="canonical-5a04e0aa455a32a54654819010bb5c767be6b3f29440d9cef505ab3d01847201"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route / 35bf2d2d4d21 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--reference--group-003.md#canonical-33ad9473556950b1d8eaeae633e15718a70e5415d43d3f778d28edc383a2e10a)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-003.md#canonical-b649a7e018cabe1e1f1ed42021ffe62bcaf11f70355eb01c1b21b08646f7df04)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-33ad9473556950b1d8eaeae633e15718a70e5415d43d3f778d28edc383a2e10a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7aaef8b2feac3fa3cef07120c15c791ae96b46765ec989b5900b0b52c3ae1512"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.labels — vn_config.inside_static_routes.static_route_list.custom_static_route.labels / 9c830a550d15 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- vn_config.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-610d9be4e688527a47eabb299762d250e4838b41b5878d842a54fc8e866d1272"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-39293efaf611a7c027373e69887d2a7bdd0b223d9402151f61e03d7b1891002c"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.labels / 9c830a550d15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e189a82bcc7f688ae7b3f7b006e634993bde966e9938b0a751cba3fb8a875d9a"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.labels / 9c830a550d15 / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0677a8f4909ff96e963ea3f4ada62de92a04c8889ab192087760ac2f64f964e"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 948fe14b1aa6 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-1a5378682ba319746be5cfbd653c6626a4c5ba467a22fd725a77a5e19f5d8a28"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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

Terraform syntax:

```terraform
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-d618c5b902f9ab86dc9957826ff6cc6dbe8b5139679e161f10dedf6b09ad705c"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 948fe14b1aa6 / 3

- [interface](resources--aws_tgw_site--reference--group-003.md#canonical-2e5e35291d21a7b757add2ccc855fadb6f7e87aba3fceff85f35c8aa4d03ccd6): complete subsection reference.

- [nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1): complete subsection reference.

<a id="canonical-59f88511c1940c8fa795f8e6203e361f9a13167488f096dbd72bf61244f2a5d3"></a>

<a id="canonical-effaa96e7777328492d03bcd1d3855e3e0f4b009c99f8d7ea9dbac125cb3ec36"></a>

## type property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 948fe14b1aa6 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-33c2e73680d9e89cee77e2bd9dc749fe81370002c74c720bcfe32fb470db8430"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 948fe14b1aa6 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--reference--group-003.md#canonical-2e5e35291d21a7b757add2ccc855fadb6f7e87aba3fceff85f35c8aa4d03ccd6)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-2e5e35291d21a7b757add2ccc855fadb6f7e87aba3fceff85f35c8aa4d03ccd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2c8b084daab5f8c6b47ec711365d1b76baa0ce1ef5a3a89afecf73c0f8de3ff"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-c442b7d2c1e6a739aa08398a51eb6169af8e23191f60666a95ee2c1bd7b6de45"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-cd6a409f7130050b3df981e158e69bd844db9259fe7cf1c27c827d9506603bf0"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 3

<a id="canonical-cc0170fdea027451204c999f65449366aa6a9720f33c0e1bb797b7a39c5e3f92"></a>

<a id="canonical-d10d8ed7a97c3a279818dc8e67f3bd44ba8762fc0a3dad09999988963b710045"></a>

## kind property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-409ced29e2a7d0ad5c3a0f8f060ffd3534e2ea34935dff79b70ab6ec89781ad9"></a>

<a id="canonical-ef512fbf6010e396d333702dec30cb65cb9427bd9c424df2fafde8d898fe8eb9"></a>

## name property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-202bc5bcabe4198cec8dada5e13538496d823d23f83dc8779ca00a4757724ba7"></a>

<a id="canonical-d977c87cbe2ba65d6d40ffd9ec341a24eefa1ec637c5aaca4b4f4fe42d2a5b46"></a>

## namespace property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
  }
}
```

<a id="canonical-cd4b8565bce845d3facb25c085ebf37b5d44b326b01ba4b13e41d21fb6d730f1"></a>

<a id="canonical-075ff188d6a222a7133ffebc24defe60cc2567fe1d6f25134a2b8140f36ae625"></a>

## tenant property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-e9150d09a3fcc16036cf3a8a6e4758c245190348a8dc598372b839b16cb2708e"></a>

<a id="canonical-909a12e18b60bd2cc985f4643801e92e3a24a1bb3693de30bb7ece7fdc400c31"></a>

## uid property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-46dbb46b833204f4532835a5159302cab9672107f8f7a8efd4e06541d97466e3"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 28d093511bbf / 9

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74b15e04e73b8a018a126c93da19e3410b1af21840054c8c0d5b27455f128dc6"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 3e7e941536e5 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-6fdde452222857aaf8b5591fb0a5bc6fe84a178ee96748079c9bcce287f1622d"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ea7e47cba4115e06ed656c1a4ae854be9e8a664cd20a077622f9abbd93b03e4"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 3e7e941536e5 / 3

- [dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-b03cba0487343972b06c3d5a2c8f4f8809a482dd53a6e7ccd8a9cfb42153e25e): complete subsection reference.

- [ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-3f7c8cfe7b161b06b5cd112ee59d1af5a16e6bb7ef3089446ff82b894cd18466): complete subsection reference.

- [ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-74f2d3409c1980645b2f885fab3179854cfeed695258b9f8ea9564c27b3a519a): complete subsection reference.

<a id="canonical-4d98e53248f7a013a58bd85dc9d38acef69b3eb7230020425a4b9b16acda873d"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 3e7e941536e5 / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-b03cba0487343972b06c3d5a2c8f4f8809a482dd53a6e7ccd8a9cfb42153e25e)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-3f7c8cfe7b161b06b5cd112ee59d1af5a16e6bb7ef3089446ff82b894cd18466)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-74f2d3409c1980645b2f885fab3179854cfeed695258b9f8ea9564c27b3a519a)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b03cba0487343972b06c3d5a2c8f4f8809a482dd53a6e7ccd8a9cfb42153e25e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-994a4a11cd184fa2fc860f7c5360fc7c2ee702d30fa3987512400858fcd41f07"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 52a7f632d302 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-b06fb45f3fbd999f7b657fb0201a9b299e8dc2545d5ba97f7aafff1e76c178e5"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-04573737f7548276e8aee5e146bd732ba118396427bad1de5ba13805d16a2acd"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 52a7f632d302 / 3

- [ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-9b3a3a7abc6a072e619c0679ef25fb1acc245afa618b55f406eadf2dabc84d38): complete subsection reference.

- [ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-0ac4d052a44849ff88a54ca0dc2ba1fb3068a1fc7a72b801102598ebeb05f902): complete subsection reference.

<a id="canonical-82ee6f09fdf4739b4361c8c27d46fd01ac6f500bbe1bc4eef52e632643c41611"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 52a7f632d302 / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-9b3a3a7abc6a072e619c0679ef25fb1acc245afa618b55f406eadf2dabc84d38)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-0ac4d052a44849ff88a54ca0dc2ba1fb3068a1fc7a72b801102598ebeb05f902)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-9b3a3a7abc6a072e619c0679ef25fb1acc245afa618b55f406eadf2dabc84d38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b53ae04b40515f7d167742b60ef4f75fd4f64ae9393068cafa77721fbe02662"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 2eb0ff3c1564 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-b03cba0487343972b06c3d5a2c8f4f8809a482dd53a6e7ccd8a9cfb42153e25e)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-fa98f7e17063309afd487927be49b9a72b18596a8ef629f336f48cd261b21d97"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-4eedd10d7d4a455928c7cd5d17944525962cdc1e0bd1ca57042a263641a67b2a"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 2eb0ff3c1564 / 3

<a id="canonical-950a7b17480e5e5da9a6acfafc3bae9976975acb9b3f1f3b48ea46e56412af7d"></a>

<a id="canonical-38301a10005eefc62a2584a91f141747f4eb6fcabf024e13778823e57001b2d6"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 2eb0ff3c1564 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1f14da1f6443054ee100f8f706e1c1e532fece7bb00837eef9e9604eb5e9a975"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 2eb0ff3c1564 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-b03cba0487343972b06c3d5a2c8f4f8809a482dd53a6e7ccd8a9cfb42153e25e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-0ac4d052a44849ff88a54ca0dc2ba1fb3068a1fc7a72b801102598ebeb05f902"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df16e004c292ddd1498bd4104340541b87ca8b49a0b600d143e086a6e603d657"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ee7d28be4224 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-b03cba0487343972b06c3d5a2c8f4f8809a482dd53a6e7ccd8a9cfb42153e25e)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-6677f3ad5914da54fa17c8fc7c4f984e6d8ca31468552a5b067b343bebb319e9"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c6e4b3c431ad76eedfa64321d567afc82e55a7383477a380503da096ff97958"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ee7d28be4224 / 3

<a id="canonical-2347ace543598f690061418e3114ca656abd2835aef1d4141498e536fde05555"></a>

<a id="canonical-4f3d19b52f7c6a170a0089d18ccf67e3680a9c87770e27451fcd209a5baa1e26"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ee7d28be4224 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-6ed5e7d92d38327c546cc93b2a27002057592d544c1b1f91b722d69dee68957b"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ee7d28be4224 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-b03cba0487343972b06c3d5a2c8f4f8809a482dd53a6e7ccd8a9cfb42153e25e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-3f7c8cfe7b161b06b5cd112ee59d1af5a16e6bb7ef3089446ff82b894cd18466"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7b92839ac9f8441684ad273d06b2cf9e1c71b48fa31303b917734005741e87f"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 629bf8d67a90 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-4badde14e648b4191aa9707927d54195111403f3b0d1fab305e45be7325e6353"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-17207631c4ec5a4b448a63da9f5ee8376539c579769274c823ebb4e0f00c5917"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 629bf8d67a90 / 3

<a id="canonical-ca08830ceb6ffbed250e9d2778d171926c21078bd4a80af02ac6bb38a8c1e889"></a>

<a id="canonical-411576e575b0eb5bae7d7c794310329e62220fae72fbb749108c33bb30671f98"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 629bf8d67a90 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-6927602211c30dd33381463482625167b8d44d8dacb93d6cfa2bd893e45bfd2c"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 629bf8d67a90 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-74f2d3409c1980645b2f885fab3179854cfeed695258b9f8ea9564c27b3a519a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cded34b8ab1a608e0d3aaa1c18b40528f6d251b873ad322956e373e26fae228"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / a179298d8da0 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-4ced369294cf0157d66ccda20be9eef137b4483699b819acabe9dc811d3167e0)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-46d9ae94dde00ec53495c90c2fcfbee6776bfb73a02401627885e96232c1c44a"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2791e95c9d112ee755f40beeb99d3e9e360db2928d1ed7b5b66b14e2569ef7b5"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / a179298d8da0 / 3

<a id="canonical-ad798ae0ad4ba2228a678866b1cbf19b999cabe5ae3ea83c1176f50aff80169c"></a>

<a id="canonical-7d7dc88ff93077d633dfcf28c41bb2e879b842bea7a3776ba7130cdd041e36d8"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / a179298d8da0 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-22844713dc518f531be0daf0d8edfa3477448410667f4fb174bfde0dd31cb077"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / a179298d8da0 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-3ddd317539c512469ed157617070988642ede349c553032a976f985d6a6239f1)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b649a7e018cabe1e1f1ed42021ffe62bcaf11f70355eb01c1b21b08646f7df04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-000c7275a1d92299e016d51437af5145bb980039c910b258b67a4557889d5af3"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.subnets — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets / d4462f44a691 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-f301e050e8b2a057dd0ce5e28e620dab933507eb388ce16a0b992c44a589b345"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3aa62b76592d667df32540156661de98252a8ded47eabe8757cb575c0a415703"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets / d4462f44a691 / 3

- [ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-ae20e706bd66c4bd3c6b48efe9e6156c2e616824adfb34f7ce9d07bee32d000b): complete subsection reference.

- [ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-edd715b210f495de208a27484febd1aa5e4e31c16a0ee2c1961df0c165fdff10): complete subsection reference.

<a id="canonical-39012738288fd4e707ca2dc88fbe8187e58194d46a397d17ef332b2419f2e676"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets / d4462f44a691 / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-ae20e706bd66c4bd3c6b48efe9e6156c2e616824adfb34f7ce9d07bee32d000b)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-edd715b210f495de208a27484febd1aa5e4e31c16a0ee2c1961df0c165fdff10)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-ae20e706bd66c4bd3c6b48efe9e6156c2e616824adfb34f7ce9d07bee32d000b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e90f5486186a8ca9267fdbcefbe714bc084829786d8ffc32774a2d64f3c7ded"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 55cc2b882aac / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-003.md#canonical-b649a7e018cabe1e1f1ed42021ffe62bcaf11f70355eb01c1b21b08646f7df04)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-2c6b9bdcd7ed16bc708705eda94c8cb275e0f7484aa0f5f38f874e4499e0f09d"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-beb5a1eb545410d622b6cf659423de49b300099a7ba67eef7d07b36cb4487b0c"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 55cc2b882aac / 3

<a id="canonical-fa671e65d6fe8bfbb3fcffd84eea55b564a30c6a0c62b9ea7b8f8c7104ec52b9"></a>

<a id="canonical-e16c98111614146ee04abbc4548a2b49b8af1ad1eeeb7dc7899b29345e43092d"></a>

## plen property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 55cc2b882aac / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-7cf6bd6a144249fe1ee478d11e0a5fcd2284984b88679adcc552e91eae6ed3fc"></a>

<a id="canonical-dca5c08a4b0b08eaac99d50e01d440e257a97010d44c23a7b3fbac4c5fac0520"></a>

## prefix property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 55cc2b882aac / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0c564a2ef73fb34c355cd07bd2b94a37b08c14872cdd19e863c317c14eec7708"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 55cc2b882aac / 6

- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-003.md#canonical-b649a7e018cabe1e1f1ed42021ffe62bcaf11f70355eb01c1b21b08646f7df04)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-edd715b210f495de208a27484febd1aa5e4e31c16a0ee2c1961df0c165fdff10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78ef95f8b5cc4aa09c39f582aa58dd0a8a2a56bab8a448057d6c3ac78c064d23"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 7aa30311bb12 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-684b125391b3d83f7858272bfa26ee42b0e6ac0e75e8967f709d45601a56ad9c)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-558b50a310c2b1bd439cae5b9e841fcfaf1758d845c3cb6c5edeab36af3e78ed)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-8e8011757c212ea8b37d25d4c558abc6d3a80681418cd52c10d598e5fcfeb1e8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-003.md#canonical-b649a7e018cabe1e1f1ed42021ffe62bcaf11f70355eb01c1b21b08646f7df04)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-d8c496090d0285aedc798d5a8ea8da20df42e86dedebaf554b72656909ad1043"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-b7fb68e9d3556decc2c3c7a14a7b46881ea1712460ce027648899a633e043f37"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 7aa30311bb12 / 3

<a id="canonical-d426422583334276659c5d9f22a0cee49e81a401c096f546c4b5ce7998ae82b1"></a>

<a id="canonical-f73b91773f9f305c125b6e78410093125f3a4fb21e8066ee998dc23ef91a5ed8"></a>

## plen property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 7aa30311bb12 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-7f1ef5e41197fc018f3679821dd6ccbd30d9ac9aaba962148e4ca2864503f934"></a>

<a id="canonical-057fd94c3ff17344289b539ade7011fbe86c993eb0d2cdc25805c89621a2052b"></a>

## prefix property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 7aa30311bb12 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-7067bccc52f4b780496713e81c33cf817f97cab4ec69a7154598ef5ae356981f"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 7aa30311bb12 / 6

- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-003.md#canonical-b649a7e018cabe1e1f1ed42021ffe62bcaf11f70355eb01c1b21b08646f7df04)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-856b6b6cc7bb0db6a24803d3cca4cece48ab9d8f72abd0805e8c4dacf4124d3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30b72f1809b46f20bec513c7956dda1b70c3f7d31f76e6a920c4abac0ad07f6f"></a>

## vn_config.no_dc_cluster_group — vn_config.no_dc_cluster_group / ed8fa735dfae / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.no_dc_cluster_group

<a id="canonical-9710f656b3aef731f795ba0202544316d932aa24090d85645f324adbc0e365e2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_dc_cluster_group = {}
```

<a id="canonical-6866ecfb01ddd95c2bb53df359473e984e2762ed96060e194830be81a57cff82"></a>

## Direct properties — vn_config.no_dc_cluster_group / ed8fa735dfae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-318474a0d6b38c0d60be78c19d651d9a07254eec3d9d934b75bbb063cd6c6861"></a>

## Next pages — vn_config.no_dc_cluster_group / ed8fa735dfae / 4

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-6894bde4d974e73e62ab179f77209c20d1a7090a847a0f0d7e08f439c2631f50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e30930286b965fce71be6d1389ac7f2a4fc869bf4cb3472c5ca77b74b4bf2721"></a>

## vn_config.no_global_network — vn_config.no_global_network / 3f57641e7534 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.no_global_network

<a id="canonical-a987642426f88fce3124633f56616d897e642eb771d8c9e239bbad1c0b7be533"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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

Terraform syntax:

```terraform
no_global_network = {}
```

<a id="canonical-d0ea3db8b618d100baeacca2929c508fc71f4b141b674e7b58ddce31c8caa1bb"></a>

## Direct properties — vn_config.no_global_network / 3f57641e7534 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0f5db03cce1e34e09e463d96420b5f6c38decd34ba958fc17b6182b3d9a3d622"></a>

## Next pages — vn_config.no_global_network / 3f57641e7534 / 4

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-307b5c27a74a99b14ef331a84e64d2222da5fdb0113aedbeec7cf78fb17b8abe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dfb4400b7f0bef625483580fd61c7c1d82ea9e76adb7cafa648cc683b9c31b0"></a>

## vn_config.no_inside_static_routes — vn_config.no_inside_static_routes / 6d2f2a8adc42 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.no_inside_static_routes

<a id="canonical-e375cdada23bca0526a524b50d8350af8e96e925797f60c37d5ba7ab959740bd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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

Terraform syntax:

```terraform
no_inside_static_routes = {}
```

<a id="canonical-a6216a11236b7a9e0f58ee4f53547ad3cefab1b46040d09e8e382af6ed1ebdcb"></a>

## Direct properties — vn_config.no_inside_static_routes / 6d2f2a8adc42 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a93ec5cc7a4e0495d75cecf6f397d3c7fc97309ffe046f594df6741d4de75ea9"></a>

## Next pages — vn_config.no_inside_static_routes / 6d2f2a8adc42 / 4

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-19f38b9fb951ac9d47b805ec6640ed9487e8fc9463758825e27e609b3268288b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf5c871ac1f806957d789bd096b68ce331f08fbfee08041a73ca3bf444221d3d"></a>

## vn_config.no_outside_static_routes — vn_config.no_outside_static_routes / 31bff73bfc84 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.no_outside_static_routes

<a id="canonical-c2df28f72af3d73a7a97eeaa0310c808eeda847d095e385bc98235a4d5a4fe6a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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

Terraform syntax:

```terraform
no_outside_static_routes = {}
```

<a id="canonical-34fb7b579eb528eb36528b5f5a5c212576eb7abf721118fc8b51a19c23e58c47"></a>

## Direct properties — vn_config.no_outside_static_routes / 31bff73bfc84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f3313cb0044fc9a30103ac4dcddd97b5dd1d20c7b899c88b32e692d04346b085"></a>

## Next pages — vn_config.no_outside_static_routes / 31bff73bfc84 / 4

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4906480550e05b5bd9fc046327e96bc2bd4eb70ec48b71cbc6737325cc2d67c6"></a>

## vn_config.outside_static_routes — vn_config.outside_static_routes / b37696a05611 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.outside_static_routes

<a id="canonical-ae75c78654be04a04bb78c2e12956605a360a6da1059caa087e36869721ee7bc"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
```

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

Terraform syntax:

```terraform
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-be456fc1c3f95beb3934de296b94752109656e0bb27f39cd2bc9aa262612c945"></a>

## Direct properties — vn_config.outside_static_routes / b37696a05611 / 3

- [static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb): complete subsection reference.

<a id="canonical-4367f916bba53e1850cbbe9aebe984b710bc877d661ff0d4195e1bad1b919fe7"></a>

## Next pages — vn_config.outside_static_routes / b37696a05611 / 4

- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ccae4be129febf060f28358f8faa56e5c6c989629912cabbae3ce3c0dce2e8f"></a>

## vn_config.outside_static_routes.static_route_list — vn_config.outside_static_routes.static_route_list / 15a04a36cb4c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- vn_config.outside_static_routes.static_route_list

<a id="canonical-2574ced6230122eec924c58971ef5a92ae3973866a065cf35d88a6d3c54dd9cb"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-ee1e6a8b7ddbd06288f9e83d185b4aab98862214047cecab221c3c7bbffe9be9"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list / 15a04a36cb4c / 3

- [custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34): complete subsection reference.

<a id="canonical-e8a1ff8dfedac86e4affc2ad0abc5bd31b3150e158ead546be5dbbcf20b54e89"></a>

<a id="canonical-7d3f72afa586393bb116999aeb9bef44847b0166a1ad1bf7ce7bc4603e455a9e"></a>

## simple_static_route property — vn_config.outside_static_routes.static_route_list / 15a04a36cb4c / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0a86b7574c901564a563c93b2644c0129efe8322e7907a24666e111027815dda"></a>

## Next pages — vn_config.outside_static_routes.static_route_list / 15a04a36cb4c / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64230a0bcbe2ffd7d1905ff7f556d0a918a740c4f0f704eb1b1e0bff47ad13d2"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route — vn_config.outside_static_routes.static_route_list.custom_static_route / 7190f306549d / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- vn_config.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-c99557f3aecc77b94fa554eff001598cee0aa1431943086473f1a9bcb2133f44"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
```

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

Terraform syntax:

```terraform
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-026c27b362ac5bfefeb2718d3ce750bfe381ee4d5d4d761f066a0b1390c6d4f2"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route / 7190f306549d / 3

<a id="canonical-0356635cc6e94e850a41a0613ff6d3bc7eeabfc9e29fd04c674cd88d44ea9bf7"></a>

<a id="canonical-54496ec6fa45ca028e2151ae3588f370ac034f0efa7b409eb8d32d5585dc445f"></a>

## attrs property — vn_config.outside_static_routes.static_route_list.custom_static_route / 7190f306549d / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_tgw_site--reference--group-003.md#canonical-2eed3842c9e0e6e9fb96c876101b22f457924467176879000bab0c0404d30861): complete subsection reference.

- [nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc): complete subsection reference.

- [subnets](resources--aws_tgw_site--reference--group-004.md#canonical-83a0fe8b56ca6eb7d2af98d9cad14496727a329df97551812930bd5ceebee4b9): complete subsection reference.

<a id="canonical-174ec3e672426acb5191b7b2ea48c1073fb37a43a3c651074dc7a150c0f39182"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route / 7190f306549d / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_tgw_site--reference--group-003.md#canonical-2eed3842c9e0e6e9fb96c876101b22f457924467176879000bab0c0404d30861)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-83a0fe8b56ca6eb7d2af98d9cad14496727a329df97551812930bd5ceebee4b9)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-2eed3842c9e0e6e9fb96c876101b22f457924467176879000bab0c0404d30861"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa23be2f43297d45d8dda8b374c4e1ac370028d9503c6ff80a088c2b2420d281"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.labels — vn_config.outside_static_routes.static_route_list.custom_static_route.labels / cfe2f89c0d3d / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- vn_config.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-7eab1046c2093f6e8a2c4830fed40c59809444adebe8a97688035749f0431aa4"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-daf2be58d38972fe665df5e31b8380d4303ba9eef141a95b5957ff57577c76e5"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.labels / cfe2f89c0d3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d783a36443249f1f82c612a9c5f96e36a795fb9d8b987ab9e4c67946a78c1cc"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.labels / cfe2f89c0d3d / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d10eb25a88a0cf8328255c77d2c2af05a95783085bf0d1e85b8d22f7371540de"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / 1bebef5c28ae / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-f929093866295e582cc6ddb53980e2feba810b8f03bc399dc0a66c5cbfe8459d"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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

Terraform syntax:

```terraform
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-da87f1b0d7a24d683e20270a5ea1ed13f23a787a46438ed4d5953859cf85e879"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / 1bebef5c28ae / 3

- [interface](resources--aws_tgw_site--reference--group-003.md#canonical-8b1a48a28a98351a4163c33558d58a78d4814063f493e0e6a758866bfb48dec0): complete subsection reference.

- [nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31): complete subsection reference.

<a id="canonical-5a2fb32d22e4ff942f51852df057470036d84734236d8a953bcc9650979a3ea3"></a>

<a id="canonical-acfb681a469b5b6a132942f79f1262435776be6d5aca7273ba1d487da739e4bf"></a>

## type property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / 1bebef5c28ae / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0fca0b11f36274e22da2d29d81779058a0d40b1f0bbb3885fe2744216d44e9c6"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / 1bebef5c28ae / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_tgw_site--reference--group-003.md#canonical-8b1a48a28a98351a4163c33558d58a78d4814063f493e0e6a758866bfb48dec0)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-8b1a48a28a98351a4163c33558d58a78d4814063f493e0e6a758866bfb48dec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-153860f89a828d5b18b0c2452531120499ddb0b6fc7941bfb2e1c26e3db13191"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-19fb2219bd74230573241ab403a66ce7e539a420a60ec015a680c1b408609715"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-b6cf0b50ebc466c9d535dfed0472cebe34fd454fa658e4ed80e49184bd7ce36b"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 3

<a id="canonical-e90bbdfe1bf893fd4d319b5ee61f13528a408b7fbe3c1021cb2c033690a0d98f"></a>

<a id="canonical-fbbe5bfbc01a7b2eec720257ff48b4ef1b88d0da9b792781b26817f6f5e9cc35"></a>

## kind property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0c1fa5b5b610f34dcf2f315a21051f9fa7d3615ee190d22a66c34c252f12851e"></a>

<a id="canonical-d3f673edfb2a713aa83ad6d306fc79084a4d8e5190d260c1eb9d85ecc5e369d5"></a>

## name property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3c2991b34c5f19b034e9ef8cf785b02a2803416fc496917f993f23bfec25face"></a>

<a id="canonical-08c98d12b71f76f9b7fbf46b228bf8cba4aa24160757c7acdc33bc66504f4d97"></a>

## namespace property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
  }
}
```

<a id="canonical-42847b715c59efd84e8f90e51be2555503d449880c4d2a2b0dbf1f3e6f1fb4d1"></a>

<a id="canonical-10c24525ff369d38a5006d70a586937f7c4a1f6f24516b81ae38309e0dc25c7b"></a>

## tenant property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3b86f7c8f475cf5d9d5eb5f315659db2a0bb1a16517787fc07d8f5f98b54b3ee"></a>

<a id="canonical-9aca52a1bda25e0d82c8e5ec5e7f687a7d5cd90640575fe9275051b5e3a3a0bd"></a>

## uid property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-fdab2572504ac0352559dd5c613863e311aea85ef6975ceba7c0e3723ea58395"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / da024daad59f / 9

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ac8f71b0dd59aa7cbf00fd58be16eac18e05f5fe540b5a48f4811725ab2ba11"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / b23dae4c1d35 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-e5ba286e4dc29ebefe41341ab7ef302cf9733e34198d930a71827a8f180d1f4e"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-144521ffd8623ac574dc10ab4b5ed3bbb3aa15505836aa2b5df97a438275fd4c"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / b23dae4c1d35 / 3

- [dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-c070a89c6db76d7dc65616a865bd4ac47cc00c547a5e24bbea96e3956f0b8d7b): complete subsection reference.

- [ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-c6a2894511543837d3125280522fda7cb4b2e23cea880f5088674675cdf88258): complete subsection reference.

- [ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-129b66eebd1df5455cc9f75880f1115a88cd8166ce565a65026960329f85f9b5): complete subsection reference.

<a id="canonical-957eb28f5d615995c0d19bcf8b6027a02410f1b359df8e62afeb7cdb36c1bb42"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / b23dae4c1d35 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-c070a89c6db76d7dc65616a865bd4ac47cc00c547a5e24bbea96e3956f0b8d7b)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-c6a2894511543837d3125280522fda7cb4b2e23cea880f5088674675cdf88258)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-129b66eebd1df5455cc9f75880f1115a88cd8166ce565a65026960329f85f9b5)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-c070a89c6db76d7dc65616a865bd4ac47cc00c547a5e24bbea96e3956f0b8d7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-219726fb117e3f2952aeae07c7367ef357f845e2cb39924b80d532ad03e3a75a"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 4a62869ec025 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-84cce9475c09d37e14b426fb590cda20187317b6689b3a125ff30973b6877454"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-e05f7147a822580e5ff37cefa4831053b5ed37adec8f9f492f2d58c971cc1166"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 4a62869ec025 / 3

- [ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-b1463096a12a58f1d79f46bbfda1ef6f044ee96cb3a9d3f10249b70dda5beb2f): complete subsection reference.

- [ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-a097b500d36e267b52e39f791b1e453915943219e205c563f994c0914ff239eb): complete subsection reference.

<a id="canonical-9764a7bf80ffb73741183cfd094359f326083736e9cbcbc72fbc441a47126a32"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 4a62869ec025 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_tgw_site--reference--group-003.md#canonical-b1463096a12a58f1d79f46bbfda1ef6f044ee96cb3a9d3f10249b70dda5beb2f)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_tgw_site--reference--group-003.md#canonical-a097b500d36e267b52e39f791b1e453915943219e205c563f994c0914ff239eb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-b1463096a12a58f1d79f46bbfda1ef6f044ee96cb3a9d3f10249b70dda5beb2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4226a656b3278444d8ac386529464eafd237f94091c581ccc27dbd1bf3fe1ca"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 94a424eb6ec1 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-c070a89c6db76d7dc65616a865bd4ac47cc00c547a5e24bbea96e3956f0b8d7b)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-1afa34227b64774358f4c6e24f04b388649bb7c04ddecfbce973f212b30fe780"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb49fec28add308612b1e279ff55e622d57e91efb0a19adc41c9026b9fc0a3b8"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 94a424eb6ec1 / 3

<a id="canonical-1e330e171479ef688aca97d6f3d87738be6f6c62329b59c0b30ec147d4be57d7"></a>

<a id="canonical-0a69628c325794140ae7da2f974b2f2ece4d264de2cec9802822e21cb10cd79d"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 94a424eb6ec1 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-f48f8537d79fd89faa07364d4ba4a56af5b59d558c3c0d664d4ede863c3cedb7"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 94a424eb6ec1 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-c070a89c6db76d7dc65616a865bd4ac47cc00c547a5e24bbea96e3956f0b8d7b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-a097b500d36e267b52e39f791b1e453915943219e205c563f994c0914ff239eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ce9b4c1573c4248464f7374b93ada0457393c29453a7d56df7c83c9987df70c"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 64da5c270bf0 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-c070a89c6db76d7dc65616a865bd4ac47cc00c547a5e24bbea96e3956f0b8d7b)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-d327fbd44591de809be627924f7b66c17ef88823d75ae436a2824cb51c936e25"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-66d6f4213c130823c3860cadf529c94ba2b0bce875f57f6c0477d8e86a7bb5b8"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 64da5c270bf0 / 3

<a id="canonical-c0aa458f3714b91a916a4bf352eeb257c4111282b25e1515dc216c6c88c7638c"></a>

<a id="canonical-4a8624f17ea90a6ac5d4065f8d84aca682113378f28133242c9a76a735ff1202"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 64da5c270bf0 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-e3d3e1321ef15e21f135e4e32631ae023d45e59725fd3b5d442f01bf38de6f01"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 64da5c270bf0 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-003.md#canonical-c070a89c6db76d7dc65616a865bd4ac47cc00c547a5e24bbea96e3956f0b8d7b)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-c6a2894511543837d3125280522fda7cb4b2e23cea880f5088674675cdf88258"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032da60cef499381e75f3710ae38c41b3247923135326920483abfa6aad91e0"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 2bb3944364ff / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-0a4b8f0d71bc907a1f25c361776aa1c65f5e0074b8ca8c8bf67ceae32f6a91a6"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-f984db4ebc1a849e04ceec717de157f90de21b90d165e474c0e95f9dabf731d4"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 2bb3944364ff / 3

<a id="canonical-c63983ec42e1d46c32d1437a2a6d53f45c17f64fbd941e3b89aeb13e607f01d5"></a>

<a id="canonical-ff0e05d38aeaa8bf00b43d631b80c503d496aa685da34f6dab1ea28098b3689b"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 2bb3944364ff / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-34b314720ed256df8309593867325e2a864e29b5815924b214c4618bf6b4b073"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 2bb3944364ff / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-129b66eebd1df5455cc9f75880f1115a88cd8166ce565a65026960329f85f9b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
