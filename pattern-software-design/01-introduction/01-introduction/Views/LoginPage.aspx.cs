using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace _01_introduction.Views
{
    public partial class WebForm1 : System.Web.UI.Page
    {
        protected void LoginBtn_Click(object sender, EventArgs e)
        {
            // username > 3 character and password cannot be empty
            string username = UsernameTB.Text;
            string password = PasswordTB.Text;

            if (username.Length <= 3)
            {
                ErrorMsg.Text = "Username must be longer than 3 characters.";
            }
            else if (string.IsNullOrEmpty(password))
            {
                ErrorMsg.Text = "Password cannot be empty.";
            }
            else
            {
                ErrorMsg.ForeColor = System.Drawing.Color.Green;
                ErrorMsg.Text = "Login successful!";
            }
        }
    }
}