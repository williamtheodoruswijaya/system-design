using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace _01_introduction.Views
{
    public partial class RegisterPage : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void RegisterBtn_Click(object sender, EventArgs e)
        {
            string username = UsernameTB.Text;
            string password = PasswordTB.Text;
            DateTime dob = DobCal.SelectedDate;
            string gender = GenderRBL.SelectedValue;
            string status = StatusDL.SelectedValue;

            if (username.Length <= 3)
            {
                errorMsg.Text = "Username must be greater than 3 character";
            } else if (string.IsNullOrEmpty(password))
            {
                errorMsg.Text = "Password cannot be empty";
            } else if (dob == DateTime.MinValue)
            {
                errorMsg.Text = "Date of Birth must be selected";
            } else if (string.IsNullOrEmpty(gender))
            {
                errorMsg.Text = "Gender cannot be empty";
            } else if (string.IsNullOrEmpty(status))
            {
                errorMsg.Text = "Status cannot be empty";
            } else
            {
                errorMsg.ForeColor = System.Drawing.Color.Green;
                errorMsg.Text = "Registration successful!";
            }
        }
    }
}