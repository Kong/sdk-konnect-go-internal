# RsaJwk

An RSA JSON Web Key (JWK) as defined in RFC 7517


## Fields

| Field              | Type               | Required           | Description        | Example            |
| ------------------ | ------------------ | ------------------ | ------------------ | ------------------ |
| `Kty`              | `string`           | :heavy_check_mark: | Key type           | RSA                |
| `Use`              | `string`           | :heavy_check_mark: | Public key use     | sig                |
| `Kid`              | `string`           | :heavy_check_mark: | Key ID             | abc123def456       |
| `Alg`              | `string`           | :heavy_check_mark: | Algorithm          | RS256              |
| `N`                | `string`           | :heavy_check_mark: | RSA modulus        |                    |
| `E`                | `string`           | :heavy_check_mark: | RSA exponent       | AQAB               |