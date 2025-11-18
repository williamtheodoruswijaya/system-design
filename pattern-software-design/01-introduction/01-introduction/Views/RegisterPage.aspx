<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="RegisterPage.aspx.cs" Inherits="_01_introduction.Views.RegisterPage" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            <h1>Register Page</h1>
            <asp:Label ID="Label1" runat="server" Text="Username"></asp:Label><br />
            <asp:TextBox ID="UsernameTB" runat="server"></asp:TextBox><br />
            <asp:Label ID="Label2" runat="server" Text="Password"></asp:Label><br/>
            <asp:TextBox ID="PasswordTB" runat="server" TextMode="Password"></asp:TextBox><br />
            <asp:Label ID="Label3" runat="server" Text="Date of Birth"></asp:Label><br />
            <asp:Calendar ID="DobCal" runat="server"></asp:Calendar><br />
            <asp:Label ID="Label4" runat="server" Text="Gender"></asp:Label><br />
            <asp:RadioButtonList ID="GenderRBL" runat="server">
                <asp:ListItem Value="Male" Selected="True">Male</asp:ListItem>
                <asp:ListItem Value="Female">Female</asp:ListItem>
            </asp:RadioButtonList>
            <asp:Label ID="Label5" runat="server" Text="Status"></asp:Label><br />
            <asp:DropDownList ID="StatusDL" runat="server">
                <asp:ListItem Value="Married" Selected="True">Married</asp:ListItem>
                <asp:ListItem Value="Single">Single</asp:ListItem>
            </asp:DropDownList><br />
            <asp:Button ID="RegisterBtn" runat="server" Text="Register" onClick="RegisterBtn_Click"/><br />
            <asp:Label ID="errorMsg" runat="server" Text=" " ForeColor="Red"></asp:Label>
        </div>
    </form>
</body>
</html>
