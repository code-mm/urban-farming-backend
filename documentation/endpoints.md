# Endpoints

## Path /device

### Path /
Allowed HTTP method: GET
Allowed content type: NONE
Required fields: NONE
Optional fields: NONE
Return status type: 200 on success or 401, 500 on error
Return content type: application/json
Description: Returns details of logged in device.

### Path /datapoint/ph
Allowed HTTP method: GET, POST
Allowed content type: application/json
Required fields: Value, Time
Optional fields: NONE
Return status type: 201 on success or 401, 500 on error
Return content type: application/json
Description: Returns either a list of all ph values for the logged in device or takes new ph values. The time field expects ISO8601 format (YYYY-MM-DDTHH:MM:SS+HH:MM).

### Path /datapoint/oxygen
Allowed HTTP method: GET, POST
Allowed content type: application/json
Required fields: Value, Time
Optional fields: NONE
Return status type: 201 on success or 401, 500 on error
Return content type: application/json
Description: Returns either a list of all oxygen values for the logged in device or takes new oxygen values. The time field expects ISO8601 format (YYYY-MM-DDTHH:MM:SS+HH:MM).

### Path /datapoint/temperature
Allowed HTTP method: GET, POST
Allowed content type: application/json
Required fields: Value, Time
Optional fields: NONE
Return status type: 201 on success or 401, 500 on error
Return content type: application/json
Description: Returns either a list of all temperature values for the logged in device or takes new temperature values. The time field expects ISO8601 format (YYYY-MM-DDTHH:MM:SS+HH:MM).

## Path /authentication

### Path /getoken/device
Allowed HTTP method: POST
Allowed content type: application/x-www-form-urlencoded
Required fields: deviceIdentifier, deviceSecret
Optional fields: NONE
Return status type: 200 on success or 401 if login credentials are wrong
Return content type: application/jwt
Description: Takes device login credentials as form data and returns JSON web token on success.