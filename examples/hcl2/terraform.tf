resource "aws_security_group_rule" "my-rule" {
  type        = "ingress"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_alb_listener" "my-alb-listener" {
  port     = "80"
  protocol = "HTTP"
  tags = {
    yor_trace = "ad78d657-8f8c-45a2-b59e-9a6915fc0aaf"
  }
}

resource "aws_db_security_group" "my-group" {

  tags = {
    yor_trace = "7a2c27aa-71bf-4b49-8cf6-1ef776b525e2"
  }
}

resource "azurerm_managed_disk" "source" {
  encryption_settings {
    enabled = false
  }
  tags = {
    yor_trace = "138d12d4-ba72-49eb-b647-3697eca46335"
  }
}
