resource "aws_cognito_user_pool" "pool" {
    name = var.COGNITO_USER_POOL

    account_recovery_setting {
        recovery_mechanism {
            name     = "verified_email"
            priority = 1
        }
    }
    
    password_policy {
        minimum_length = 8
        require_lowercase = true
        require_numbers = true
        require_symbols = true
        require_uppercase = true
    }

    schema {
        name = "name"
        attribute_data_type = "String"
        mutable = true
        required = true
    }

    schema {
        name = "email"
        attribute_data_type = "String"
        mutable = true
        required = true
    }

    schema {
        name = "phone_number"
        attribute_data_type = "String"
        mutable = true
        required = true
    }

    schema {
        name = "api_token"
        attribute_data_type = "String"
        mutable = false
        required = false
    }

    tags = {
        Environment = "PROD"
        Name        = "COGNITO AUTH"
    }

}