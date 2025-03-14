<h1 align="center">Password Cracker</h1>

<h2 align="center">Overview</h2>
<p>Password Cracker is a Go-based tool designed to recover passwords by employing two primary attack methods:</p>
<ul>
  <li><strong>Brute-Force Attack</strong>: Systematically attempts all possible combinations of characters up to a specified length.</li>
  <li><strong>Dictionary Attack</strong>: Utilizes a predefined list of potential passwords to identify the correct one.</li>
</ul>

<h2 align="center">Features</h2>
<ul>
  <li><strong>Multiple Hashing Algorithms</strong>: Supports MD5, SHA-1, and SHA-256 hashing functions.</li>
  <li><strong>Customizable Charset</strong>: Allows users to define the character set for brute-force attacks.</li>
  <li><strong>User-Friendly Interface</strong>: Prompts users for input and guides them through selecting the desired attack method.</li>
</ul>

<h2 align="center">Usage</h2>
<ol>
  <li><strong>Clone the Repository</strong>:
    <pre><code>git clone https://github.com/anonymByte-404/password-cracker.git
cd password-cracker</code></pre>
  </li>
  <li><strong>Run the Application</strong>:
    <pre><code>go run cmd/main.go</code></pre>
    <p>The program will prompt you to:</p>
    <ul>
      <li><strong>Enter the Password</strong>: Type the password you wish to crack. The application will hash this input using the selected hashing algorithm.</li>
      <li><strong>Choose an Attack Type</strong>: Select either <code>brute-force</code> or <code>dictionary</code>.</li>
    </ul>
    <p>Based on your selection, the program will attempt to crack the password using the chosen method.</p>
  </li>
</ol>

<h2 align="center">Attack Methods</h2>

<h3>Brute-Force Attack</h3>
<p>The brute-force attack method systematically tries all possible combinations of characters up to a specified length. You can customize the character set and maximum password length as needed.</p>

<h3>Dictionary Attack</h3>
<p>The dictionary attack method uses a predefined list of potential passwords (a dictionary file) to find a match. Ensure you have a suitable dictionary file (e.g., <code>dictionary.txt</code>) placed in the project directory.</p>

<h2 align="center">Hashing Algorithms</h2>
<p>The application supports the following hashing algorithms:</p>
<ul>
  <li><strong>MD5</strong>: A widely used hash function producing a 128-bit hash value.</li>
  <li><strong>SHA-1</strong>: Produces a 160-bit hash value, commonly used in various security applications.</li>
  <li><strong>SHA-256</strong>: Part of the SHA-2 family, producing a 256-bit hash value.</li>
</ul>
<p>You can modify the hashing function in the code as needed.</p>

<h2 align="center">License</h2>
<p align="center">This project is licensed under the MIT License - see the <a href="LICENSE">LICENSE</a> file for details.</p>
