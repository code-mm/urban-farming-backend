# Authentication of devices against the device endpoint

## Request of JSON web token
The device requires a valid device identifier and device secret to successfully receive a JSON web token.
The identifier consists of an UUID and the secret is an alphanumeric string, special characters are permitted.
To start the authentication, the device has to send its identifier and secret to the "/authentication/gettoken/device" endpoint.
The endpoint requires the HTTP method POST and the content type "application/x-www-form-urlencoded". The required form fields
are called deviceIdentifier and deviceSecret. A possible request could look like "deviceIdentifier=<identifier>&deviceSecret=<secret>",
while the <identifier> and <secret> is replaced with its respective values.
If the authentication was succesful the endpoint responds with HTTP status code 200 and the content type "application/jwt",
otherwise a HTTP status code 401 is send.

## Authentication against the device endpoint
After a successful acquisition of a JSON web token, it must be send with each subsequent request.
Therefore, the HTTP header "Authorization" must be added and the value must be of form "Bearer <JSON web token>",
while <JSON web token> is replaced with the actual token.

The validity of a token is restricted to one hour. After this period a new token must be acquired.