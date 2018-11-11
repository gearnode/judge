```
grpcurl -d '{
  "policy": {
    "name": "Demo Policy",
    "description": "Test policy for create policy test",
    "document": {
      "version": "v1alpha1",
      "statements": [
        {
          "effect": "Allow",
          "actions": ["judge:CreatePolicy"],
          "resources": ["orn:judge-org:policy-service::foo/bar"]
        }
      ]
    }
  }
}' 127.0.0.1:5053 judge.api.v1alpha1.Judge.CreatePolicy
```

