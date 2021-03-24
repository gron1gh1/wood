import { Button, FormControl,FormLabel ,Input} from "@chakra-ui/react"

function App() {
  return (
    <div style={{display:'flex',justifyContent:'center',alignItems:'center',height:'100vh'}}>
    <FormControl id='login' style={{width:'500px'}}>
      <FormLabel>Email</FormLabel>
      <Input type='email' placeholder='Email'></Input> 
      <FormLabel>Password</FormLabel>
      <Input type='password' placeholder='Password'></Input> 
      <Button style={{marginTop:'10px'}}width="500px" colorScheme="teal" type="submit">Login</Button>
    </FormControl>
    </div>
  );
}

export default App;
