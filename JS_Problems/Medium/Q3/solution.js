function updateEmployeeSalary(employees, employeeId, newSalary) {
    // Iterate through the array to find the employee by ID
    for (let i = 0; i < employees.length; i++) {
        if (employees[i].id === employeeId) {
            // Update the salary of the found employee
            employees[i].salary = newSalary;

            // Return the updated employee object
            return employees[i];
        }
    }

    // Return null if no employee with the given ID was found
    return null;
}
