---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-3030111232031013-1211000020302020-2211333321002330-3230023012301230-1001332213202202-0023331323233113-2211201212323102-1020201123011313"></a>

#### `tls_parameters.cert_params.cipher_suites` property

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2010133011013011-2321002333310321-2022113303113202-3300320203000233-0022230011331101-1000101201031201-0130202123221321-1203301012101331"></a>

<a id="canonical-0123121323211322-1133231023112133-1100321013131021-2332331022310003-1311031112221012-3022001231311110-0012022013013332-0313101131222332"></a>

#### `tls_parameters.cert_params.maximum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1202221312032002-1021120001313223-0232231313211033-3222022132031310-0001221231221101-0110131030310012-1231233013023032-3100032120103000"></a>

<a id="canonical-0033000211313311-0320220111203321-1311210232221123-2030120013101001-2300201331332103-2203003022132223-2330310113021203-1233200300023103"></a>

#### `tls_parameters.cert_params.minimum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [skip_server_verification](data-sources--cluster--reference--group-002.md#canonical-2203313332122113-1223012210100202-0131312102322332-3013232002313323-3211011121201321-1332231132202223-2312120011100112-0311122122233210): complete subsection reference.

- [tls_validation_params](data-sources--cluster--reference--group-002.md#canonical-0311330302130000-1103112110213333-2011012322122311-2233023231023333-0223121323131323-0321101110002022-2023211210232300-0200112202213030): complete subsection reference.

- [volterra_trusted_ca](data-sources--cluster--reference--group-002.md#canonical-3021330033200220-1330310122233123-3231102223321030-3023222322301122-1122321031021203-2321223232312101-1111010311110033-3031011212332212): complete subsection reference.

<a id="canonical-3200233330322300-0222311312232113-3213320003030123-1322013221032212-2322301321223233-0222313020122231-3003230032101323-3022130211123332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.certificates` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002)
- tls_parameters.cert_params.certificates

<a id="canonical-3022330102202210-3123330220312113-0312132232002130-1111231323023032-1232011120320001-3203220210233132-1311120232000002-2322322001002210"></a>

Type: `"list"`. Computed.

Client TLS Certificate required for mTLS authentication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3221201021123030-1223211023230020-1113123322221100-2312110310221110-0033230111302202-2010030222131332-1312313030000302-2323001330211322"></a>

### Direct properties for `tls_parameters.cert_params.certificates`

<a id="canonical-1222332013333121-2013213300323201-3210202202232113-0020001322233133-3122312021202221-0321331012012121-2000200203331112-1003030013203130"></a>

#### `tls_parameters.cert_params.certificates.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2121132212133111-3300221213100133-3313320231300011-2001030012221111-3223333311123321-0002001101012333-1113301321100002-0101033220033121"></a>

<a id="canonical-0022231221033232-2000221313200210-2320220121211222-2211000123103011-2231001223333021-1222321300311331-3022311321013213-0013322322302332"></a>

#### `tls_parameters.cert_params.certificates.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2223313113220012-2131123100102022-0131321023130333-2020102211013302-0210000233201101-3020213302210212-0111003330312130-1003033312323332"></a>

<a id="canonical-0323200200131331-2103020311120211-0211310131013111-1302030223000030-1322301000201223-1200012030210302-2033100011323033-3101332010023023"></a>

#### `tls_parameters.cert_params.certificates.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0102331323120132-3001112132113332-1211230012011312-3120233121113221-2031020132220023-2123332221023220-2310221113301203-1120330201123310"></a>

<a id="canonical-1322302112122333-3303033320203213-1333321121332301-3330131133221031-3313211211120100-2003310122213010-1013032001231132-3022103013112000"></a>

#### `tls_parameters.cert_params.certificates.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1103111223313121-2020323202333033-2122012003311331-0233130213110020-3000011021332102-3031012233103103-2312111222303223-2133000222233200"></a>

<a id="canonical-3212233233003102-1231031033011122-3323031013120011-3332312320210100-3100131230132201-2031111220212003-1211132113022202-2323002333131112"></a>

#### `tls_parameters.cert_params.certificates.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2203313332122113-1223012210100202-0131312102322332-3013232002313323-3211011121201321-1332231132202223-2312120011100112-0311122122233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.skip_server_verification` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002)
- tls_parameters.cert_params.skip_server_verification

<a id="canonical-2301303311111120-3032310021011233-2003021203003001-3111220333323103-3331102010030210-1303331012011200-1320023131012332-2122220003321302"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311330302130000-1103112110213333-2011012322122311-2233023231023333-0223121323131323-0321101110002022-2023211210232300-0200112202213030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.tls_validation_params` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002)
- tls_parameters.cert_params.tls_validation_params

<a id="canonical-1011101011210320-2100123233010303-2121233123002021-1101123103313012-1101212103212131-3231311102101010-1233301013312121-1302013320113100"></a>

Type: `"single"`. Computed.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-2302102000011211-3020311301212220-1122312000230313-0032112223110221-3032011020332001-3130103210133113-3002101220222222-2021332132022321"></a>

### Direct properties for `tls_parameters.cert_params.tls_validation_params`

<a id="canonical-1313312222133121-0100233022131132-3113213130111102-0101212020011311-0231001120201130-1103022132231011-2021001211031332-3013021200011022"></a>

#### `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` property

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](data-sources--cluster--reference--group-002.md#canonical-2032231233123332-0323221312003131-3320301013001022-0321230320311212-2213110311020101-3201312031112203-1033313033111330-2312131313120033): complete subsection reference.

<a id="canonical-1103100231322000-1131303332003023-2312202300102333-3333222330311022-3131113100130312-1322020033133212-2322311310003020-2311213012212332"></a>

<a id="canonical-1013121010210133-3311210132122201-0132322212103311-2211133300320011-3222320203320122-1111102131020323-1301322210211132-3320012112130032"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2323113302302121-3201200330103312-1003132131303332-0221210221313110-2203002300223012-0113003002203330-2201011121330013-2330130320020110"></a>

<a id="canonical-3210301210332230-3113000101023201-3203112233311303-2230022322232321-0302112102301322-1020313030302020-0020102200130021-1221310122311023"></a>

#### `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-2032231233123332-0323221312003131-3320301013001022-0321230320311212-2213110311020101-3201312031112203-1033313033111330-2312131313120033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.tls_validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002)
- [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-002.md#canonical-0311330302130000-1103112110213333-2011012322122311-2233023231023333-0223121323131323-0321101110002022-2023211210232300-0200112202213030)
- tls_parameters.cert_params.tls_validation_params.trusted_ca

<a id="canonical-0212023021313002-3212032010001011-3211112322220121-1132302212011220-0331022313012321-3320113321121202-1333320121011203-3020122231300122"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-3332223201232200-2130301331320013-1023103012320020-2213110321223220-3013033310201110-2122330200302230-0221210012110220-3221122310210202"></a>

### Direct properties for `tls_parameters.cert_params.tls_validation_params.trusted_ca`

- [trusted_ca_list](data-sources--cluster--reference--group-002.md#canonical-3031133302223002-3232203332213322-1212222330212302-0331110233002201-2313332331312332-3132212112232230-3010122111112223-0221120002233013): complete subsection reference.

<a id="canonical-3031133302223002-3232203332213322-1212222330212302-0331110233002201-2313332331312332-3132212112232230-3010122111112223-0221120002233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002)
- [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--reference--group-002.md#canonical-0311330302130000-1103112110213333-2011012322122311-2233023231023333-0223121323131323-0321101110002022-2023211210232300-0200112202213030)
- [tls_parameters.cert_params.tls_validation_params.trusted_ca](data-sources--cluster--reference--group-002.md#canonical-2032231233123332-0323221312003131-3320301013001022-0321230320311212-2213110311020101-3201312031112203-1033313033111330-2312131313120033)
- tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-0223323032202213-2313010202310121-0231312000012330-0001221022123103-3233103002332310-3032312121001113-2232312000333212-1033232233010232"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1213323203102101-2112130020003020-3103332012332121-0030301120201123-3022110133303221-0021213111110302-2113122020303303-2120333303033213"></a>

### Direct properties for `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-0023212330332120-2000300333233321-2220113133103233-3230003110110102-2131212020211011-3123013131103223-1103110331202122-0131023320232123"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3320100030132230-2210322320001312-1321301333103323-0323333023032021-1320123302013313-3110001202113212-3202311300101221-2210311113202330"></a>

<a id="canonical-1230313330123001-3232233013110201-1220011201112211-2110103013121003-1212301133131131-1132033212123331-1001211001102002-2313211010122030"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3121210121330303-1332330222021232-3203310320232033-3232101220331302-0002320213010213-0322303011112102-1122232131120210-1030010033013032"></a>

<a id="canonical-0200203020012330-1220003233203022-1031200110213231-2233213210132121-3102302023032002-0020110303211011-1313330212132113-1301213323332211"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1011233230030113-1200220311101223-0032122222132322-0032303321322111-0313123210021233-2212032312212023-0201000001021032-3032323333023032"></a>

<a id="canonical-1201001131113100-2221223331233032-1130101303332230-1301233011130120-0321333213212211-0210322123212313-2022130112112223-1112221100332203"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3230002123103210-2111302233121132-2132000030103011-2013223212002120-3211201023100321-1210333110030102-0002310103110021-2213211232103000"></a>

<a id="canonical-1132002230213202-0313212111022200-2123100123030313-1130033310010012-1101301303011301-2303120220111010-1132200222233112-1202002332023001"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3021330033200220-1330310122233123-3231102223321030-3023222322301122-1122321031021203-2321223232312101-1111010311110033-3031011212332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.cert_params](data-sources--cluster--reference--group-001.md#canonical-1302321232321230-3200120210232323-3303020302200300-0001013030021031-3331321221003122-2201331232132222-3133133322200212-0023331301020002)
- tls_parameters.cert_params.volterra_trusted_ca

<a id="canonical-3133123202302013-2331133120202210-3230012023023100-0131003221300200-1102310033223221-1103111021210120-1302333123331121-3330320113202330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- tls_parameters.common_params

<a id="canonical-3123321333013122-3232130001010230-3132300121003211-3311133311011300-3030121031211001-2123330211121210-0110211021032022-0301300233023302"></a>

Type: `"single"`. Computed.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

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

<a id="canonical-2230003321223230-0313030013010331-0222332102123321-1203213001010130-1200321011123013-2030101013101233-1321313200120302-1311210220111200"></a>

### Direct properties for `tls_parameters.common_params`

<a id="canonical-0313133310232211-0223111322203322-1320220230012222-1313210210022130-3102311323321320-2123030033221220-0202023332302323-1302023331311021"></a>

#### `tls_parameters.common_params.cipher_suites` property

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0211131202031123-1022120120303223-0302001332103201-0112330002030233-2122101302121030-3133011303222321-2132001302232310-0101021330031133"></a>

<a id="canonical-1333133012013112-1230031323303223-2230231032013311-0123112232323333-0032130232322102-3231132213333202-2020132022333213-2023213312313302"></a>

#### `tls_parameters.common_params.maximum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2023303003031310-3132233101231002-0330332023201012-3001311100331221-0131322303212333-0302003032031230-2112300311313030-0020331023211220"></a>

<a id="canonical-3131202123033321-0302313303100331-3331031122221001-1033233220321122-0231113112303302-1322002133333302-2311112200002212-3030100110332032"></a>

#### `tls_parameters.common_params.minimum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_certificates](data-sources--cluster--reference--group-002.md#canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303): complete subsection reference.

- [validation_params](data-sources--cluster--reference--group-002.md#canonical-0200022102031013-2222112231032311-1120133023100231-3231311323232033-1210120332312132-3201303000300002-2210000030022320-3323232202330311): complete subsection reference.

<a id="canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- tls_parameters.common_params.tls_certificates

<a id="canonical-1310312112200001-0111003312110310-1221220032000223-2333122121202220-1203203223012011-2001320231031300-2231012200331331-2313021222330210"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

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

<a id="canonical-0133221103223131-0123333101102333-2203112310131303-2302332002223221-0131323211323331-2332210132122133-2330230223122120-0001300333313202"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates`

<a id="canonical-2121031032221002-0001010332023210-0321021211233032-3010201330102002-3122033311001131-0320231312030111-2022013333223232-0030332331222120"></a>

#### `tls_parameters.common_params.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--cluster--reference--group-002.md#canonical-3211220231023030-0211211010003022-2032221100030203-1100230020220002-1301212132013212-0310232100030232-3122001203330322-2033323332203220): complete subsection reference.

<a id="canonical-2022200303201210-2233203331313022-1023133203113332-0332123313301210-0333220313210100-0010330212001211-3202331220332313-3302133301123013"></a>

<a id="canonical-1233211213202210-1020220322020011-2311202332223333-3310300103102222-3021023033132303-3312113230131032-1212003211312220-0031122322221123"></a>

#### `tls_parameters.common_params.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--cluster--reference--group-002.md#canonical-0333022001133111-1013013311312210-0021031102333020-0002303122221100-2221233021011100-1213212022120331-0101230132331312-2123112121220223): complete subsection reference.

- [private_key](data-sources--cluster--reference--group-002.md#canonical-3332021020211332-3123012023131012-2101111123031023-1031202112323121-2230300020231100-0213301131133110-2112212023220111-3001032021203232): complete subsection reference.

- [use_system_defaults](data-sources--cluster--reference--group-002.md#canonical-1203310001013021-3231300320223222-0211200123111301-1231202010331201-1232233003300232-1233203300231030-3310210002133233-2331301221313323): complete subsection reference.

<a id="canonical-3211220231023030-0211211010003022-2032221100030203-1100230020220002-1301212132013212-0310232100030232-3122001203330322-2033323332203220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-002.md#canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-1022021001021210-1102331000011200-1200021100011011-1102323022131112-1002301221112013-1133301311010301-0002212332201330-0332303233211221"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-1203113210231131-0232133230131232-3333002202330123-2122010132020311-0123230111212110-0232001010101013-2112020221203221-0113111303321233"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-0100111200330331-2312131231313103-1310302203321222-2003210301302323-3131323222310322-1013101302122133-1000312102330230-2000301031030321"></a>

#### `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333022001133111-1013013311312210-0021031102333020-0002303122221100-2221233021011100-1213212022120331-0101230132331312-2123112121220223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-002.md#canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-0230131330101331-1213223131333112-3133012001313320-1013210110100232-3022231301003022-3312023202020021-3023213001222123-0223221010111103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332021020211332-3123012023131012-2101111123031023-1031202112323121-2230300020231100-0213301131133110-2112212023220111-3001032021203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-002.md#canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-1220301332013011-3020203200102211-3312112101033032-3103030112303233-0311101121210230-1211123102010021-2303033320012333-1102010033231300"></a>

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

<a id="canonical-2222300210311020-0333001230211022-0331101112002131-0333022033003302-1112111332122222-3331133010313212-0131210223310032-0303331022211220"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--cluster--reference--group-002.md#canonical-1021322332311313-3220121300303322-3113202021110223-2110032102110021-3032233210110213-0301130133011032-3121133111211033-0333022022130101): complete subsection reference.

- [clear_secret_info](data-sources--cluster--reference--group-002.md#canonical-3233013132112032-0102123333110330-0131101210123131-3022030213033012-0032000130322210-1322000001311310-1102231301030203-0302130200232102): complete subsection reference.

<a id="canonical-1021322332311313-3220121300303322-3113202021110223-2110032102110021-3032233210110213-0301130133011032-3121133111211033-0333022022130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-002.md#canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-002.md#canonical-3332021020211332-3123012023131012-2101111123031023-1031202112323121-2230300020231100-0213301131133110-2112212023220111-3001032021203232)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0303003102332031-3102312333322202-1322121222033023-2230122310203002-2210103101301032-0201332313223121-3301030032100001-1011202311332000"></a>

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

<a id="canonical-2320321332212123-3112021021110032-1211122203012311-1120001031103110-0223233201031121-2210100000120001-1023223321221002-2013102121123211"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0300321200000013-0321130202121000-3030103031000123-0000003031131302-2103132113230330-0220012300121113-1300202323112200-0111130111311100"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1110111221123001-1213120331130003-1302023330123001-3232302110132230-2232303320003201-1010210021311012-3033220122032233-3212111231020130"></a>

<a id="canonical-1102020300233212-0112101200303202-0113230023331321-2330323123202020-1332102112112311-2021301222212201-1223311110331212-2031302210333031"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3222103111231110-2001133301203021-1320022022111032-3121230302302130-3210232113112200-2230202000113013-2002201131221020-3112231012301201"></a>

<a id="canonical-3121301311022220-3211312311213020-0332333011131100-3233121232132010-1011002233130230-3001002212023212-0203023302231322-1032321233311123"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3233013132112032-0102123333110330-0131101210123131-3022030213033012-0032000130322210-1322000001311310-1102231301030203-0302130200232102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-002.md#canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--reference--group-002.md#canonical-3332021020211332-3123012023131012-2101111123031023-1031202112323121-2230300020231100-0213301131133110-2112212023220111-3001032021203232)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-3230133002223220-3322210003133223-1001231001012213-3232233120213013-1002022110121103-3302223001333103-3223333213123103-2231021121202102"></a>

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

<a id="canonical-2000210221320030-3332130210030320-3030100133031212-0300230020102032-1110313100011121-2011230210323020-0230321103210211-0023133011332332"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2321233323321113-0133220021200321-0021123133303033-1123230220310130-2232023323212321-3031023001132112-3303221203302203-3111032101231213"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2313231013332230-1101332021330321-0330213313031132-3230302302022231-2203122221101022-2203233031222023-3121201212102311-3213310131320120"></a>

<a id="canonical-2030133233111332-0111023012030223-1213103323303100-2123132120010023-1311220303332311-0013331301002120-1121330130232002-0003200010132221"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1203310001013021-3231300320223222-0211200123111301-1231202010331201-1232233003300232-1233203300231030-3310210002133233-2331301221313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--reference--group-002.md#canonical-2322321211020111-2323120130310233-0122031120101221-3000012023333102-2330032330010331-3200222000200001-3020003201231013-1001330012223303)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-0222220232131013-1232322220032302-3231213130201123-3102220102201223-3131022302133303-1111131220021330-1211010321211221-3130303132223220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200022102031013-2222112231032311-1120133023100231-3231311323232033-1210120332312132-3201303000300002-2210000030022320-3323232202330311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- tls_parameters.common_params.validation_params

<a id="canonical-0322110113201121-1223001312302130-3221220020333311-2102121120311333-1201312300212232-1001021232023032-2212201223012011-0202313332100230"></a>

Type: `"single"`. Computed.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-1313211312001021-3103321221120203-3322122301321331-2203210210013111-1123010220233311-2221302000221021-1113212310103003-1120111030331233"></a>

### Direct properties for `tls_parameters.common_params.validation_params`

<a id="canonical-2013330002022110-3221233320111022-2113200032223021-2131100333310012-2022101303332013-2112121220032300-1011101133030330-3012122112003203"></a>

#### `tls_parameters.common_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](data-sources--cluster--reference--group-002.md#canonical-2131030032102121-2232130121132133-2222233320102200-3020321232333101-3320331002022311-0230030310122103-1122032113030020-1201030230302231): complete subsection reference.

<a id="canonical-2203102031001022-1232331020313232-2201021032232100-1323211222313033-3110112333211222-0132100313332010-0013133333323313-0223133231332311"></a>

<a id="canonical-2001310303310111-0312013310022132-1211002231202331-1002321101211130-1313131323331332-2223301221112333-1111213201121000-1223322102312131"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0113323122033110-3033312012301320-3322210232211311-0302120230123012-2100011100011203-3201201302002033-3231030021101332-3112220022000330"></a>

<a id="canonical-2102310233211013-0300222223323132-0212333000122232-2233010002212203-0203033232030302-0211003111121303-3310021310131131-2112233121021110"></a>

#### `tls_parameters.common_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-2131030032102121-2232130121132133-2222233320102200-3020321232333101-3320331002022311-0230030310122103-1122032113030020-1201030230302231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-002.md#canonical-0200022102031013-2222112231032311-1120133023100231-3231311323232033-1210120332312132-3201303000300002-2210000030022320-3323232202330311)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-1001022012100032-2110022203120003-3001223300010213-1210330122130203-2311223123122123-1012010020130323-3333000331332001-0131012011101101"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-1002133022001311-0030033012020301-2102133313012101-3022033001032002-1021223222323230-0023331100133123-1202221120121031-2033312110303330"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca`

- [trusted_ca_list](data-sources--cluster--reference--group-002.md#canonical-1302023231010103-1211132201113330-2120003000221312-1102232100121322-1332322200210011-0312002010201213-0022230120021133-3103112121230330): complete subsection reference.

<a id="canonical-1302023231010103-1211132201113330-2120003000221312-1102232100121322-1332322200210011-0312002010201213-0022230120021133-3103112121230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- [tls_parameters.common_params](data-sources--cluster--reference--group-002.md#canonical-2020301003100311-1020232013200323-1320133212303201-3201233100312213-2211323232011020-0212102222130302-3231321120323312-1213210113111313)
- [tls_parameters.common_params.validation_params](data-sources--cluster--reference--group-002.md#canonical-0200022102031013-2222112231032311-1120133023100231-3231311323232033-1210120332312132-3201303000300002-2210000030022320-3323232202330311)
- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--cluster--reference--group-002.md#canonical-2131030032102121-2232130121132133-2222233320102200-3020321232333101-3320331002022311-0230030310122103-1122032113030020-1201030230302231)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-3112213033112310-3321000031223300-0223001133022332-1233122302331310-0302101111221020-1031000022220032-2000002220333113-3121322311102031"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3301121021030133-1110103303210213-3303033212320130-3011021333331233-2221233212110213-1203202033223111-1200133323320211-1121332213330231"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-1202320330213302-1023021011101132-1031201131232203-0200212202022221-2303100011202131-1313012103132123-0313333100213330-0233123120110002"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3030211113230023-3231233232030313-0010220100023131-1200102300012110-1222022112312002-1320323323101021-2211211212332212-3003320032330332"></a>

<a id="canonical-1310001332020321-1032223310100031-2001113112021300-1012231121312322-1010320300222330-1131132011030333-2313333003001203-2310220233332311"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3213131101002101-0031213131020203-0202222210230021-0023223100103303-2030331232022202-1103002203120120-2321000011213130-0331330013132003"></a>

<a id="canonical-3012003001000102-3122231133021222-2102330303213301-1303130201033131-0232132033313202-2320133020300202-0131230230202331-0023000313303311"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0301131130020121-1102321302221000-2213122213100100-0110031033331033-1112123312020313-3322220012302102-0231132301303121-0011101311131123"></a>

<a id="canonical-0321212223001233-0221111203310313-1022233112320310-1101331121331321-1032003310011023-3333121021220030-0110020032233111-3120332303231321"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2323321121011121-1323102303012131-2300113320111023-0222203320220011-1232120212213323-1013320011330013-0301031012032103-3323322032211210"></a>

<a id="canonical-1130212001211333-0033203332131212-3000011323203200-3203011210313221-0022010013222032-0133010321211222-0122203131332020-0013323313202031"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0030303201213312-2002233020110002-0103213003202313-0231023111222112-0213131221031201-1302110222212321-3200100310301122-3232332322003330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- tls_parameters.default_session_key_caching

<a id="canonical-0020132322020000-2220201301121210-3130113132111221-2123311202220002-1010313220223102-0220312132310131-1321201312200331-3130110101312233"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131003313100023-2011121223333200-0211332321102113-3302131121032312-1301111301003330-3020313203010010-1213131322111210-1232103311201020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- tls_parameters.disable_session_key_caching

<a id="canonical-0013102131022122-0322331031333310-1133301322200323-3130000112013123-1313212320200100-1023012212332300-1100212021011323-0010230221220121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130130133200210-1010212021311333-2111122320031101-0333233313010021-0323213330130012-2122211020012332-0212311203233110-0313123311210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.disable_sni` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- tls_parameters.disable_sni

<a id="canonical-2223002332030032-1012200132302111-3202003203333003-0230030132030210-0221333123123002-2313312130233001-1121312003002202-2132113112132320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223202311322031-0130323103302030-3212303312011321-1130013200232211-2120320001110303-0210223212222320-3120323210123323-3231021123203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-3000331110120102-1313102013232321-0131210230113302-0010210120302330-0120211233030302-1220001101121231-0213222102031020-1101021100303031)
- tls_parameters.use_host_header_as_sni

<a id="canonical-2013110333111220-1222102320022320-3221310000202032-1012033010233330-1133110033133110-3122322311111221-0311121001312330-2100213320030111"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313001211201301-0220311332132120-1031132020332300-1331101230303030-2110211232132032-3032132311122313-0201333030230221-2110023012110012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- upstream_conn_pool_reuse_type

<a id="canonical-0022220233122111-0100101301121003-2023000303213331-1302120002102111-1012311112211200-0101133111202031-1133222103113333-1221211022000220"></a>

Type: `"single"`. Computed.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

<a id="canonical-0302233320231212-2220133221023322-2312322231130203-2101321211011021-0311333100111312-3232002233012000-1332013003130310-3021120332213202"></a>

### Direct properties for `upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-3100330022133011-1311223112002301-0000310020011311-0302201301230001-1211102111320303-2021102023222112-2021130211020312-3233303110123010): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-1201023120022202-1133133111310330-1211223333003113-2200301021101033-1131032220121222-1330020213203012-2321322033220003-1021132220223303): complete subsection reference.

<a id="canonical-3100330022133011-1311223112002301-0000310020011311-0302201301230001-1211102111320303-2021102023222112-2021130211020312-3233303110123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-0313001211201301-0220311332132120-1031132020332300-1331101230303030-2110211232132032-3032132311122313-0201333030230221-2110023012110012)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-0211233012013322-3003020020021130-3312103330213112-3233101213002130-2131003200102012-3211331012023103-2112300133111102-0211203032320210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable conn pool reuse.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201023120022202-1133133111310330-1211223333003113-2200301021101033-1131032220121222-1330020213203012-2321322033220003-1021132220223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-2020311230233013-2301231333200031-0000133110133310-2200103231211020-2210231033130123-0133012030321013-2220221320010211-0023211323121200)
- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-0313001211201301-0220311332132120-1031132020332300-1331101230303030-2110211232132032-3032132311122313-0201333030230221-2110023012110012)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-1111133231100310-0321301220320333-2112112032010331-3320133000332222-3020122220111331-3301322223110012-1130221203102232-2223220032131301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable conn pool reuse.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
