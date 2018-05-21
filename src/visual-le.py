#-*- encoding: utf-8 -*
import wx

"""
by 王紫璇
日志的时间好多重合的，直接按照文件顺序来了
自动1s刷新一次
"""

class MainFrame(wx.Frame):
	
	def __init__(self):
		wx.Frame.__init__(self, None, -1, "drawer-pudding", (0, 0), (1200, 800))
		self.Bind(wx.EVT_PAINT, self.OnPaint)
		self.timer = wx.Timer(self)#创建定时器
		self.Bind(wx.EVT_TIMER, self.OnTimer, self.timer)#绑定一个定时器事件
		self.timer.Start(1000)#设定时间间隔
		self.Center(wx.BOTH) # 窗口居中显示
		self.pos = (0,0)
		self.pen = wx.Pen("green", 3)
		self.brush = wx.Brush('', wx.TRANSPARENT)  
		#日志里的两段都在condition里了
		self.condition = [['STATE','0','peer 0','follower',''],['STATE','0','peer 2','follower',''],['STATE','0','peer 1','follower',''],['TESTCASE','Initial Election'],['STATE','0','peer 2','follower','elec timeout'],['STATE','1','peer 2','candicate',''],['MESSG','0','peer 1 <- peer 2','Rcv ReqVote(1)'],['MESSG','1','peer 2 <- peer 1','ReqVote Rep ,(1, true)'],['STATE','1','peer 2','candicate','voted sucess'],['STATE','1','peer 2','leader',''],['MESSG','0','peer 0 <- peer 2','Rcv ReqVote(1)'],['MESSG','1','peer 1 <- peer 2','Rcv ApdEntry(1)'],['MESSG','1','peer 1 <- peer 2','Rcv ApdEntry(1)'],['MESSG','1','peer 2 <- peer 1','ApdEntry Rep, (1, true)'],['MESSG','1','peer 2 <- peer 0','ReqVote Rep ,(1, true)'],['STATE','1','peer 1','follower','newTerm'],['STATE','1','peer 1','follower',''],['STATE','1','peer 0','follower','newTerm'],['STATE','1','peer 0','follower',''],['MESSG','1','peer 0 <- peer 2','Rcv ApdEntry(1)'],['MESSG','1','peer 2 <- peer 1','ApdEntry Rep, (1, true)'],['MESSG','1','peer 0 <- peer 2','Rcv ApdEntry(1)'],['MESSG','1','peer 2 <- peer 0','ApdEntry Rep, (1, true)'],['MESSG','1','peer 2 <- peer 0','ApdEntry Rep, (1, true)'],['MESSG','1','peer 1 <- peer 2','Rcv ApdEntry(1)'],['MESSG','1','peer 2 <- peer 1','ApdEntry Rep, (1, true)'],['MESSG','1','peer 0 <- peer 2','Rcv ApdEntry(1)'],['MESSG','1','peer 2 <- peer 0','ApdEntry Rep, (1, true)'],['TESTCASE','Disconnect 2'],['STATE','1','peer 0','follower','elec timeout'],['STATE','2','peer 0','candicate',''],['MESSG','1','peer 1 <- peer 0','Rcv ReqVote(2)'],['MESSG','2','peer 0 <- peer 1','ReqVote Rep ,(2, true)'],['STATE','2','peer 0','candicate','voted sucess'],['STATE','2','peer 0','leader',''],['STATE','2','peer 1','follower','newTerm'],['STATE','2','peer 1','follower',''],['MESSG','2','peer 1 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (2, true)'],['MESSG','2','peer 1 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (2, true)'],['TESTCASE','Connect 2'],['MESSG','1','peer 2 <- peer 1','ApdEntry Rep, (0, false)'],['MESSG','1','peer 2 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 1 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (2, true)'],['MESSG','2','peer 0 <- peer 2','ApdEntry Rep, (2, true)'],['STATE','2','peer 2','leader','new Term'],['STATE','2','peer 2','follower',''],['MESSG','2','peer 2 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 2','ApdEntry Rep, (2, true)'],['MESSG','2','peer 1 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (2, true)'],['MESSG','2','peer 2 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 2','ApdEntry Rep, (2, true)'],['MESSG','2','peer 1 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (2, true)'],['MESSG','2','peer 1 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (2, true)'],['MESSG','2','peer 2 <- peer 0','Rcv ApdEntry(2)'],['MESSG','2','peer 0 <- peer 2','ApdEntry Rep, (2, true)'],['TESTCASE','Disconnect 0,1'],['STATE','2','peer 1','follower','elec timeout'],['STATE','3','peer 1','candicate',''],['STATE','2','peer 2','follower','elec timeout'],['STATE','3','peer 2','candicate',''],['STATE','3','peer 1','candicate','elect timeout'],['STATE','4','peer 1','candicate',''],['STATE','3','peer 2','candicate','elect timeout'],['STATE','4','peer 2','candicate',''],['STATE','4','peer 1','candicate','elect timeout'],['STATE','5','peer 1','candicate',''],['STATE','4','peer 2','candicate','elect timeout'],['STATE','5','peer 2','candicate',''],['STATE','5','peer 1','candicate','elect timeout'],['STATE','6','peer 1','candicate',''],['STATE','5','peer 2','candicate','elect timeout'],['STATE','6','peer 2','candicate',''],['STATE','6','peer 1','candicate','elect timeout'],['STATE','7','peer 1','candicate',''],['TESTCASE','Connect 1'],['STATE','6','peer 2','candicate','elect timeout'],['STATE','7','peer 2','candicate',''],['MESSG','7','peer 1 <- peer 2','Rcv ReqVote(7)'],['MESSG','7','peer 2 <- peer 1','ReqVote Rep ,(7, false)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (0, false)'],['MESSG','7','peer 2 <- peer 0','ReqVote Rep ,(0, false)'],['MESSG','2','peer 0 <- peer 1','ApdEntry Rep, (0, false)'],['MESSG','7','peer 1 <- peer 0','ReqVote Rep ,(0, false)'],['STATE','7','peer 1','candicate','elect timeout'],['STATE','8','peer 1','candicate',''],['MESSG','7','peer 2 <- peer 1','Rcv ReqVote(8)'],['STATE','8','peer 2','candicate','newTerm'],['STATE','8','peer 2','follower',''],['MESSG','8','peer 1 <- peer 2','ReqVote Rep ,(8, true)'],['STATE','8','peer 1','candicate','voted sucess'],['STATE','8','peer 1','leader',''],['MESSG','8','peer 2 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 2','ApdEntry Rep, (8, true)'],['MESSG','8','peer 2 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 2','ApdEntry Rep, (8, true)'],['MESSG','8','peer 2 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 2','ApdEntry Rep, (8, true)'],['MESSG','2','peer 0 <- peer 2','ApdEntry Rep, (0, false)'],['MESSG','8','peer 2 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 2','ApdEntry Rep, (8, true)'],['TESTCASE','Connect 0'],['MESSG','2','peer 0 <- peer 2','ApdEntry Rep, (0, false)'],['MESSG','8','peer 2 <- peer 0','ReqVote Rep ,(0, false)'],['MESSG','2','peer 0 <- peer 2','Rcv ReqVote(3)'],['MESSG','8','peer 2 <- peer 0','ReqVote Rep ,(3, true)'],['STATE','3','peer 0','leader','new Term'],['STATE','3','peer 0','follower',''],['MESSG','3','peer 0 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 2 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 2','ApdEntry Rep, (8, true)'],['MESSG','8','peer 1 <- peer 0','ApdEntry Rep, (8, true)'],['STATE','8','peer 0','follower','newTerm'],['STATE','8','peer 0','follower',''],['MESSG','8','peer 0 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 0','ApdEntry Rep, (8, true)'],['MESSG','8','peer 2 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 2','ApdEntry Rep, (8, true)'],['MESSG','8','peer 1 <- peer 0','ReqVote Rep ,(0, false)'],['MESSG','8','peer 0 <- peer 1','Rcv ReqVote(5)'],['MESSG','8','peer 1 <- peer 0','ReqVote Rep ,(8, false)'],['MESSG','8','peer 2 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 2','ApdEntry Rep, (8, true)'],['MESSG','8','peer 0 <- peer 1','Rcv ApdEntry(8)'],['MESSG','8','peer 1 <- peer 0','ApdEntry Rep, (8, true)'],['MESSG','8','peer 0 <- peer 2','ApdEntry Rep, (0, false)'],['MESSG','8','peer 0 <- peer 1','ApdEntry Rep, (0, false)']]
		#toumingtianchong
		self.SetBackgroundColour("white")
		self.colorA="green"
		self.colorB="green"
		self.colorC="green"
		self.testText=''
		self.i=0
		size = self.GetClientSize() 
		self.buffer = wx.Bitmap(size.width, size.height)
		dc = wx.BufferedDC(None, self.buffer)
		dc.SetPen(self.pen) 
		dc.SetBackground(wx.Brush(self.GetBackgroundColour()))
		dc.SetBrush(self.brush)
		dc.Clear()
		self.condition1=self.condition[0]
		#初始化，三个都是follower
		self.stA="follower"
		self.stB="follower"
		self.stC="follower"
		self.termA='0'
		self.termB='0'
		self.termC='0'

	def OnTimer(self, evt):#显示时间事件处理函数    
		dc = wx.BufferedDC(wx.ClientDC(self))
		self.Draw(dc)
		self.i=self.i+1
		if self.i>=len(self.condition ):
			return 0
		self.condition1=self.condition[self.i]
	def OnPaint(self, event):
		dc = wx.BufferedDC(wx.ClientDC(self))
		self.Draw(dc)
		self.i=self.i+1
		self.condition1=self.condition[self.i]
	#最左边
	def OptionsA(self, dc):
		term=''
		temp=''
		if self.condition1[0]=='STATE':
			if self.condition1[2]=='peer 0':
				self.termA=self.condition1[1]
				temp=self.condition1[4]
				if self.condition1[3]=='follower':
					self.colorA="green"
					self.stA="follower"
				if self.condition1[3]=='candidate':
					self.colorA="blue"
					self.stA="candidate"
				if self.condition1[3]=='leader':
					self.colorA="red"
					self.stA="leader"
		dc.SetPen(wx.Pen(self.colorA, 3))
		dc.SetFont(wx.Font(15, wx.SWISS, wx.NORMAL, wx.NORMAL))
		dc.DrawEllipse(100, 100, 200, 200)
		dc.DrawText('term: '+self.termA+' peer0: '+self.stA+' '+temp, 100, 70)
	#中间的
	def OptionsB(self, dc):
		term=''
		temp=''
		if self.condition1[0]=='STATE':
			if self.condition1[2]=='peer 1':
				self.termB=self.condition1[1]
				term=self.condition1[1]
				temp=self.condition1[4]
				if self.condition1[3]=='follower':
					self.colorB="green"  
					self.stB="follower"
				if self.condition1[3]=='candidate':
					self.colorB="blue"
					self.stB="candidate"
				if self.condition1[3]=='leader':
					self.colorB="red"
					self.stB="leader"
		dc.SetPen(wx.Pen(self.colorB, 3))
		dc.SetFont(wx.Font(15, wx.SWISS, wx.NORMAL, wx.NORMAL))
		dc.DrawEllipse(500, 400, 200, 200)
		dc.DrawText('term: '+self.termB +' peer1: '+self.stB+' '+temp, 500, 630)
		#右边的
	def OptionsC(self, dc):
		term=''
		temp=''
		if self.condition1[0]=='STATE':
			if self.condition1[2]=='peer 2':
				self.termC=self.condition1[1]
				temp=self.condition1[4]
				if self.condition1[3]=='follower':
					self.colorC="green"
					self.stC="follower"
				if self.condition1[3]=='candidate':
					self.colorC="blue"
					self.stC="candidate"
				if self.condition1[3]=='leader':
					self.colorC="red"
					self.stC="leader"
		dc.SetPen(wx.Pen(self.colorC, 3))
		dc.SetFont(wx.Font(15, wx.SWISS, wx.NORMAL, wx.NORMAL))
		dc.DrawEllipse(900, 100, 200, 200)
		dc.DrawText('term: '+self.termC + ' peer2: '+self.stC+' '+temp, 900, 70)
	#如果是msg，调用
	def OptionsMsg(self, dc):
		if self.condition1[0]=='MESSG':
			term=self.condition1[1]
			dc.SetPen(wx.Pen("black", 1))
			dc.SetFont(wx.Font(15, wx.SWISS, wx.NORMAL, wx.NORMAL))
			if self.condition1[2]=='peer 0 <- peer 2':
				dc.DrawRotatedText('peer 0 <-- peer 2', 400, 180, 0)
				dc.DrawRotatedText(self.condition1[3], 400, 210, 0)
				dc.DrawLine(300,200,900,200)
			if self.condition1[2]=='peer 2 <- peer 0':
				dc.DrawRotatedText('peer 0 --> peer 2', 400, 180, 0)
				dc.DrawRotatedText(self.condition1[3], 400, 210, 0)
				dc.DrawLine(300,200,900,200)
			if self.condition1[2]=='peer 2 <- peer 1':
				dc.DrawRotatedText('peer 1 --> peer 2', 750, 400, 37)
				dc.DrawRotatedText(self.condition1[3], 770, 420, 37)
				dc.DrawLine(680,440,920,260)
			if self.condition1[2]=='peer 1 <- peer 2':
				dc.DrawRotatedText('peer 1 <-- peer 2', 750, 400, 37)
				dc.DrawRotatedText(self.condition1[3], 770, 420, 37)
				dc.DrawLine(680,440,920,260)
			if self.condition1[2]=='peer 0 <- peer 1':
				dc.DrawRotatedText('peer 0 <-- peer 1', 320, 260, -37)
				dc.DrawRotatedText(self.condition1[3], 300, 300, -37)
				dc.DrawLine(280,260,520,440)
			if self.condition1[2]=='peer 1 <- peer 0':
				dc.DrawRotatedText('peer 0 --> peer 1', 320, 260, -37)
				dc.DrawRotatedText(self.condition1[3], 300, 300, -37)
				dc.DrawLine(280,260,520,440)
	def OptionsTest(self, dc):
		if self.condition1[0]=='TESTCASE':
			self.testText=self.condition1[1]
		dc.DrawRotatedText('TestCase : '+ self.testText, 500, 20, 0)
	def Draw(self, dc):
		dc.Clear()
		dc.SetPen(wx.Pen("green", 3))
		self.OptionsA(dc)
		self.OptionsB(dc)
		self.OptionsC(dc)
		if self.condition1[0]=='MESSG':
			self.OptionsMsg(dc)
		self.OptionsTest(dc)

if __name__ =='__main__':
	app = wx.App()
	frame = MainFrame()
	frame.Show()
	app.MainLoop()
