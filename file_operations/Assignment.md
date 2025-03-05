# Assignment

File Operations in Golang

## Task 1: File Creation and Writing

**Objective**: Implement a function to create a new log file and write initial data to it.

**Instructions**:

- Write a function `CreateLogFile(fileName string, initialContent string) error` that:

  - Creates a new file with the given `fileName`.

  - Writes the `initialContent` to the file.

  - Handles errors appropriately, ensuring the file is closed after writing.

**Expected Outcome**:

- A file named `service_log.txt` is created with the initial content "Service started successfully.".

- The function should handle cases where the file cannot be created (e.g., due to permission issues) and return an appropriate error message.


## Task 2: Reading and Displaying File Content

**Objective**: Implement a function to read the contents of a log file and display it.

**Instructions**:

- Write a function `ReadLogFile(fileName string) (string, error)` that:

  - Opens the specified file.

  - Reads the entire content of the file.

  - Returns the content as a string and any errors encountered.

**Expected Outcome**:

- The content of `service_log.txt` is printed to the console.

- The function should handle cases where the file does not exist or cannot be opened and return an appropriate error message.

## Task 3: Appending to a File

**Objective**: Extend the log file by appending new entries.

**Instructions**:

- Write a function `AppendToLogFile(fileName string, contentToAdd string) error` that:

  - Opens the file in append mode.

  - Appends the new content to the file.

  - Closes the file after appending and handles any errors.

**Expected Outcome**:

- The content "New event logged at: 2024-09-04 12:00:00." is appended to `service_log.txt`.

- The function should ensure that new content is added without overwriting the existing data.

## Task 4: File Statistics Reporting

**Objective**: Analyze and report statistics about the log file.

**Instructions**:

- Write a function `ReportFileStats(fileName string) error` that:

  - Retrieves and prints the following information about the file:

    - File size in bytes.

    - Number of lines in the file.

    - Last modification time.

  - Handles errors related to file access and existence.

**Expected Outcome**:

- The tool prints out the size, line count, and last modification time of `service_log.txt`.

- The function should accurately calculate the number of lines by reading the file.


## Task 5: Implementing Error Handling and Logging

**Objective**: Improve the robustness of the tool by adding comprehensive error handling and logging.

**Instructions**:

- Modify the existing functions to ensure all possible errors are logged to a separate file called `error_log.txt`.

- Implement a function `LogError(errorMessage string)` that:

  - Opens (or creates) `error_log.txt` in append mode.

  - Write the error message with a timestamp to the file.

**Expected Outcome**:

- Any errors encountered by the tool are logged to `error_log.txt`.

- The tool should continue to function even if an error is encountered, logging it for later analysis.


## Task 6: Implementing File Pointer Navigation

**Objective**: Add the ability to read specific parts of a log file based on a file pointer.

**Instructions**:

- Write a function `ReadFromOffset(fileName string, offset int64, length int) (string, error)` that:

  - Moves the file pointer to the specified `offset`.

  - Reads `length` bytes from the file starting at that offset.

  - Returns the read content and any errors encountered.

**Expected Outcome**:

- The tool should accurately read and return the content from the specified position in the file.

- The function should handle edge cases, such as reading beyond the file’s end or starting from an invalid offset.


## Task 7: Deleting Log Files

**Objective**: Implement functionality to safely delete log files when they are no longer needed.

**Instructions**:

- Write a function `DeleteLogFile(fileName string) error` that:

  - Deletes the specified file.

  - Logs an error if the file cannot be deleted.

**Expected Outcome**:

- The specified file should be deleted from the filesystem.

- The function should ensure that the file is no longer needed before deletion and log any errors encountered.
