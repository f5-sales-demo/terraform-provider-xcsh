---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-1212033231013101-1211210110120033-2201033211111131-0322310331021203-2200012100333231-0333231210321010-3113021122010322-1233011230030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110)
- primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint

<a id="canonical-1310013002022232-3123100310320020-3302132323103132-3010312231002020-0322211133101202-0330130210120012-0023132001322001-0221131011332221"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 fingerprint.

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

<a id="canonical-3120020033303301-2121222331032011-2310132003021012-3003323130230022-3121302303222233-1013133220113302-3222110010033010-3230310132120213"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint`

<a id="canonical-0232000100001222-2300203103302233-3121333113330031-3221212301303111-2210323011131110-1130103021323111-0102312333130231-1103232131000330"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` property

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-0321003213330333-0032202123013222-1320212310111100-0222320102323300-1220220323000302-1020321111210022-3020001101001020-3133201013032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110)
- primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

<a id="canonical-2021231231131011-2002010300220312-2013311332002322-2323103231333301-3032102030233320-0321333100133023-2223003113133021-3222311112110112"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 fingerprint.

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

<a id="canonical-2203321130112220-0010121310031111-3033232202321133-1102121332021222-2333033220321011-2321101130233212-1322012101121121-0021213313210232"></a>

### Direct properties for `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint`

<a id="canonical-3201203001133212-0131133232032213-2011311202003222-2021102032301211-1201033223000230-2302323122302213-2013030113111010-0021122211020113"></a>

#### `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` property

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.tlsa_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.tlsa_record

<a id="canonical-1101223212221102-3322031333001130-1103322111122133-1033323303213231-3332201223232100-1321231100221231-2130122203112013-0222012110223001"></a>

Type: `"single"`. Computed.

Configuration parameter for tlsa record.

Additional upstream details:

DNS TLSA Record.

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

<a id="canonical-3303003102112123-2332323333222123-3030023233200231-0223112213210010-3022103311103221-3222312102031313-2010300101033320-3030212101102023"></a>

### Direct properties for `primary.rr_set_group.rr_set.tlsa_record`

<a id="canonical-3013011123320300-1330002211301302-0010222000332022-0331110332323222-3013010011022032-0310312331031133-3232332231113010-0011100130233301"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.name` property

Type: `"string"`. Computed.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-2300322000010100-3330131011030130-0002231310231000-0231233002101022-1131210322121011-2330020130222313-3123203000313102-2213110333130232): complete subsection reference.

<a id="canonical-2300322000010100-3330131011030130-0002231310231000-0231233002101022-1131210322121011-2330020130222313-3123203000313102-2213110333130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.tlsa_record.values` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121)
- primary.rr_set_group.rr_set.tlsa_record.values

<a id="canonical-2312332012210332-2122213103112133-1132301230330200-0223000200211331-1121010310320232-3020220333130203-3313313323312102-2330321021132333"></a>

Type: `"list"`. Computed.

TLSA Value. Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1232100322033103-0301321000102231-2230312332221101-3120133331222020-2223310122011200-0301111320131331-0221302033332223-3021212102320211"></a>

### Direct properties for `primary.rr_set_group.rr_set.tlsa_record.values`

<a id="canonical-2022331111110123-2100103133330003-2011211220002120-2200222331011202-0233133231121211-2100130131020002-0020133032101201-1120100312130101"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` property

Type: `"string"`. Computed.

The actual data to be matched given the settings of the other fields.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1032033121302201-0313000230112110-1201323010101122-2203310201303130-2302113130330111-3201013302201122-1023203000012300-0202130302300122"></a>

<a id="canonical-2031012221303131-3332323320112021-1203331231220032-3223222203332033-0022130020222223-0012021213131222-1100033133100203-0110210010323132"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` property

Type: `"string"`. Computed.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2231300020001110-2202032131312021-3013103200231130-1122101031311230-3002123203312100-1023230201233010-3121330111100032-2221122330021322"></a>

<a id="canonical-3000003322113333-0020313230002223-0220011110113103-2131231113012212-3133302321010321-3202110012003201-2121033100200333-3231032002121032"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` property

Type: `"string"`. Computed.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0301100322000222-2210310111030100-0120200130323031-0313300323302321-2020023000003023-2230202232131031-2302031133313230-0222231030103131"></a>

<a id="canonical-3112210202330221-3023222021312330-0022032102002303-3313302323202122-1302232030010022-3321211130301103-1030132201121300-1313300100033311"></a>

#### `primary.rr_set_group.rr_set.tlsa_record.values.selector` property

Type: `"string"`. Computed.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2112000100003211-0010122023020330-2302300331333210-3123120333012012-1313313003312213-2120120330311231-1033321003121122-0030131112120013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.rr_set_group.rr_set.txt_record` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.txt_record

<a id="canonical-3120013321330311-3132233230321201-0111310300333022-1333330321012022-3022221032332233-0201113210031120-0121111121022320-2322222212030010"></a>

Type: `"single"`. Computed.

DNSTXTResourceRecord.

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

<a id="canonical-2110221123332000-0100113202301020-3233333232112311-1233123230323030-3212130223322320-0322323120001321-0032321131231313-2011201012233331"></a>

### Direct properties for `primary.rr_set_group.rr_set.txt_record`

<a id="canonical-1230123002302333-0320100331032333-0332032213020213-2313010123113222-3130032100232102-0001031233212030-3111013300000210-2232023222320232"></a>

#### `primary.rr_set_group.rr_set.txt_record.name` property

Type: `"string"`. Computed.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2313020101101202-1330213132331132-0000110133301131-0012232110110003-3303020000303202-3130122103003331-0120203331232122-3001322231031230"></a>

<a id="canonical-2210311300233332-0021220212320021-0112110201010231-3201010033211232-3102001111113122-3012332133313202-2331210200000131-1220203123112323"></a>

#### `primary.rr_set_group.rr_set.txt_record.values` property

Type: `["list", "string"]`. Computed.

Text. Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3111301200310120-2012211111103233-2001110322021102-1132112033133212-1100110333112313-0210200023120120-3233033130322010-1032311322222221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `primary.soa_parameters` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.soa_parameters

<a id="canonical-3332200032033121-0332301033131232-0222202113000111-3302133332331313-1302101233232222-2003033200120003-1030130033320323-1003113213310120"></a>

Type: `"single"`. Computed.

Configuration parameter for soa parameters.

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

<a id="canonical-2000323312202313-3313311222210220-3311022310213231-2000013112010100-2312312330112331-0222133303131330-2302003002323012-3312201321222033"></a>

### Direct properties for `primary.soa_parameters`

<a id="canonical-1201033000303111-2223303120002331-0021033000232021-3123122112002012-1123232021130110-3023211303113021-3311312133110210-0002231111232302"></a>

#### `primary.soa_parameters.expire` property

Type: `"number"`. Computed.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3100223313012320-2031021332301022-2331201032203202-2002023020101122-2110133031121231-3201313011302023-0120221133100100-3220113013302122"></a>

<a id="canonical-1100321231311131-3102132121012021-3233333120031020-3303103130023223-2011303231111331-1312100303231312-0131332103000003-1102010312031121"></a>

#### `primary.soa_parameters.negative_ttl` property

Type: `"number"`. Computed.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3120102023112030-3232121120322003-0021302112212101-2201012221110322-0220003032222203-2310310313302020-1202332331112213-0131133131130230"></a>

<a id="canonical-0121020012321202-1003302201222223-2122020113220131-3122333230111202-1100212003322102-2132231032233003-3033301323321213-1200333101230201"></a>

#### `primary.soa_parameters.refresh` property

Type: `"number"`. Computed.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-0132113021222221-2013001303010022-3221122332130011-0033113000132220-1303321123221002-0302320220223033-2122331011201321-2323110332310111"></a>

<a id="canonical-2231012012022110-0121030302321322-3221000133200312-1131022023123301-0300333233000310-3133131233132321-2331231203033103-2132113112121022"></a>

#### `primary.soa_parameters.retry` property

Type: `"number"`. Computed.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3213120213122202-0213310131301121-3202133221121323-0102313213023020-0023231121222111-1321320022330200-1010321303222311-0122010110212101"></a>

<a id="canonical-0223101302023122-0221001203313233-1110100130123232-0121001010323032-2322123301223102-0310021020001122-0333201221111021-1103231011232322"></a>

#### `primary.soa_parameters.ttl` property

Type: `"number"`. Computed.

TTL. SOA record time to live (in seconds)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- secondary

<a id="canonical-3313001100221131-0333113010030320-1300031010002223-1231331333010010-0122201331100023-3022321102020300-0200013201131111-3320301233010323"></a>

Type: `"single"`. Computed.

SecondaryDNSCreateSpecType.

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

<a id="canonical-3011000233321030-1210221220011022-0223012133332002-1203301022213112-3103313021322001-2131210023301331-2232120020222311-3321133331111001"></a>

### Direct properties for `secondary`

<a id="canonical-2231231030010330-2312101322001132-3330301120301212-2132112131013123-3301303301023113-1220032023023021-1321133312132220-0110301231200022"></a>

#### `secondary.primary_servers` property

Type: `["list", "string"]`. Computed.

Configuration parameter for primary servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3330223310100001-0110311112003331-3013103213122000-2103321322103210-2012323002310000-1321203300233211-2200023201100210-2022222311131132"></a>

<a id="canonical-0313111032123333-1200023110023203-3332230001333222-2023223000210310-3312032012300232-0110003202103332-3301323033313203-3122032201033333"></a>

#### `secondary.tsig_key_algorithm` property

Type: `"string"`. Computed.

\[Enum: HMAC\_MD5|UNDEFINED|HMAC\_SHA1|HMAC\_SHA224|HMAC\_SHA256|HMAC\_SHA384|HMAC\_SHA512\] TSIG
key-value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC\_MD5:
HMAC\_MD5 - HMAC\_SHA1: HMAC\_SHA1 - HMAC\_SHA224: HMAC\_SHA224 - HMAC\_SHA256: HMAC\_SHA256 -
HMAC\_SHA384: HMAC\_SHA384 - HMAC\_SHA512: HMAC\_SHA512. Possible values are \`HMAC\_MD5\`,
\`UNDEFINED\`, \`HMAC\_SHA1\`, \`HMAC\_SHA224\`, \`HMAC\_SHA256\`, \`HMAC\_SHA384\`,
\`HMAC\_SHA512\`. Defaults to \`UNDEFINED\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNDEFINED",
  "enum": [
    "HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2032021232222221-3331232103113031-0222131211003030-0111213310012111-0120232201320103-3220323322013310-1223033222120332-1100022012211321"></a>

<a id="canonical-3231302330132132-2132302322010212-1022032300313220-2020110012323213-3030223300300323-0331333100311010-2302300121003000-0021223200001333"></a>

#### `secondary.tsig_key_name` property

Type: `"string"`. Computed.

TSIG key name as used in TSIG protocol extension.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331): complete subsection reference.

<a id="canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary.tsig_key_value` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- secondary.tsig_key_value

<a id="canonical-3101233323332301-0312030330010020-0112200003233111-3030302033213132-1231033132112202-1221323011331233-3132101020311220-0322033320313121"></a>

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

<a id="canonical-0122300113320102-0210103231020302-2301330303101212-0212002120201123-0301313010211222-0200331103010110-1203212031011322-0233201233020303"></a>

### Direct properties for `secondary.tsig_key_value`

- [blindfold_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-1132333021301033-3331233222032213-2213011102012210-1120300021003122-3020030301333202-2313113001320103-1101220211020023-0103200133333003): complete subsection reference.

- [clear_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-0323113331312310-3323000001233012-3330320133313030-1133231201132300-2233133131302311-2333301023012213-1023201101130113-2221321221301002): complete subsection reference.

<a id="canonical-1132333021301033-3331233222032213-2213011102012210-1120300021003122-3020030301333202-2313113001320103-1101220211020023-0103200133333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary.tsig_key_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331)
- secondary.tsig_key_value.blindfold_secret_info

<a id="canonical-3322021301003221-1011000011013122-3221210313313322-1010211101010230-0021322012320102-3301222202013311-0330313211203121-3310320323300313"></a>

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

<a id="canonical-2001210210011020-2133130230033120-0221331023212100-1223031123232031-2331201302011030-2301111231332002-0130301233222103-2330122332123102"></a>

### Direct properties for `secondary.tsig_key_value.blindfold_secret_info`

<a id="canonical-2210020032131012-1313301020210112-2223100111000300-1320110301123013-0121113322102313-0110122232020112-2301213112313002-3122000002222200"></a>

#### `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300313031010331-1211322202132302-0323221133333301-2120210322212120-2320303010232122-3302202122310030-0301313211010021-1100333111010221"></a>

<a id="canonical-0223303311032112-3221001213310101-0031103302022200-0202230312100213-2101210202203320-0200332311210320-2023222110232020-1303030311001302"></a>

#### `secondary.tsig_key_value.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1332321201022032-0211212221300332-2022310122013230-2211233030113203-0220013000023200-1133221300112003-2122103012323023-0322031022312012"></a>

<a id="canonical-1123301212202101-2222211320200032-2211311020120220-3312001131122313-3020333122331111-1100212220111300-0202121031121112-3123031130122000"></a>

#### `secondary.tsig_key_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0323113331312310-3323000001233012-3330320133313030-1133231201132300-2233133131302311-2333301023012213-1023201101130113-2221321221301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `secondary.tsig_key_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331)
- secondary.tsig_key_value.clear_secret_info

<a id="canonical-0002213001331323-1200100121010123-2301120030233222-2203330130212332-1331302213010022-1330213231323202-2010321231203310-0110223012100222"></a>

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

<a id="canonical-3301231120200213-2332023121211221-0003213122231212-3033112032213220-3023223220302203-0121301201203310-3002231220303113-2320303213120311"></a>

### Direct properties for `secondary.tsig_key_value.clear_secret_info`

<a id="canonical-3012021300301121-2020233020332311-3000200220320123-2000003303023023-3202331302310133-2231133310330133-3302321211300300-3300030030031111"></a>

#### `secondary.tsig_key_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1330123220130302-2101200121133331-3010022021113122-3331023211020110-2230132031312100-0101030030100311-3213331132110111-3103103113303113"></a>

<a id="canonical-3221021301023132-1033233313110022-0130112312131310-0322122022100033-2230330112212321-1232221210332110-0201220013310012-2013030120303000"></a>

#### `secondary.tsig_key_value.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
