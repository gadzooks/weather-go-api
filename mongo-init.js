db.createUser(
    {
        user: "integUser",
        pwd: "integPass",
        roles: [
            {
                role: "readWrite",
                db: "weatherIntegDb"
            }
        ]
    }
);