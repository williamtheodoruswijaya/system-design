<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="LoginPage.aspx.cs" Inherits="_01_introduction.Views.WebForm1" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            <h1>Login Page</h1>
            <asp:Label ID="Label1" runat="server" Text="Username"></asp:Label><br />
            <asp:TextBox ID="UsernameTB" runat="server"></asp:TextBox><br />
            <asp:Label ID="Label2" runat="server" Text="Password"></asp:Label><br/>
            <asp:TextBox ID="PasswordTB" runat="server" TextMode="Password"></asp:TextBox><br />
            <asp:CheckBox ID="RememberMeCB" runat="server" Text="Remember Me"/><br />
            <asp:Button ID="LoginBtn" runat="server" Text="Login" onClick="LoginBtn_Click"/><br />
            <asp:Label ID="ErrorMsg" runat="server" Text=" " ForeColor="Red"></asp:Label>
        </div>
    </form>
</body>
</html>
